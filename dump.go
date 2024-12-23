package main

import (
	"fmt"
	"io"
	"iter"
	"strings"
	"unicode"
)

func verbose(w *strings.Builder, bytes []byte) {
	s := strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return '.'
	}, string(bytes))

	w.WriteString(fmt.Sprintf("|%s|", s))
}

func dumpSeq(r io.Reader, ps bool, v bool, cols int, offs int) iter.Seq2[string, error] {
	format := " %02X"
	padding := 3
	if ps {
		format = "%02X"
		padding = 2
	}

	return func(yield func(string, error) bool) {
		bytes := make([]byte, cols)

		for n, err := r.Read(bytes); n != 0; n, err = r.Read(bytes) {
			if err != nil {
				yield("", err)
				return
			}

			var out strings.Builder

			if !ps {
				out.WriteString(fmt.Sprintf("%016X", offs))
			}

			for _, b := range bytes[:n] {
				out.WriteString(fmt.Sprintf(format, b))
			}
			out.WriteString(fmt.Sprintf("%*s", padding*(cols-n), ""))

			if v {
				out.WriteString("\t")
				verbose(&out, bytes[:n])
			}

			out.WriteString("\n")

			if !yield(out.String(), err) {
				break
			}

			offs += n
		}
	}
}
