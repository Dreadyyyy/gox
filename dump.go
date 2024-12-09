package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

func unixSeq(r io.Reader, verbose bool, cols int) func(func(string, error) bool) {
	return func(yield func(string, error) bool) {
		bytes := make([]byte, cols)
		offs := 0

		for n, err := r.Read(bytes); n != 0; n, err = r.Read(bytes) {
			if err != nil {
				yield("", err)
				return
			}

			var out strings.Builder

			out.WriteString(fmt.Sprintf("%016X:", offs))

			for _, b := range bytes[:n] {
				out.WriteString(fmt.Sprintf(" %02X", b))
			}

			out.WriteString(fmt.Sprintf("%*s", 3*(cols-n), ""))

			s := strings.Map(func(r rune) rune {
				if unicode.IsGraphic(r) {
					return r
				}
				return '.'
			}, string(bytes[:n]))

			if verbose {
				out.WriteString(fmt.Sprintf("\t|%s|", s))
			}

			out.WriteString("\n")

			if !yield(out.String(), err) {
				break
			}

			offs += n
		}
	}
}
