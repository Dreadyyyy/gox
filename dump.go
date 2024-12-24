package main

import (
	"fmt"
	"io"
	"os"
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

func dump(in io.Reader, out *os.File, plain bool, verb bool, cols int, offs int) {
	format := " %02X"
	padding := 3
	if plain {
		format = "%02X"
		padding = 2
	}

	buff := make([]byte, cols)

	for n, err := in.Read(buff); n != 0; n, err = in.Read(buff) {
		if err != nil {
			fatal(err.Error())
		}

		var sb strings.Builder

		if !plain {
			sb.WriteString(fmt.Sprintf("%016X", offs))
		}

		for _, b := range buff[:n] {
			sb.WriteString(fmt.Sprintf(format, b))
		}
		sb.WriteString(fmt.Sprintf("%*s", padding*(cols-n), ""))

		if verb {
			sb.WriteString("\t")
			verbose(&sb, buff[:n])
		}

		sb.WriteString("\n")

		fmt.Fprint(out, sb.String())

		offs += n
	}
}
