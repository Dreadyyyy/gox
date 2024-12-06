package main

import (
	"fmt"
	"strings"
	"unicode"
)

func unixSeq(bytes []byte, verbose bool, line int) func(func(string) bool) {
	return func(yield func(string) bool) {
		for i := 0; i < len(bytes); i += line {
			out := fmt.Sprintf("%016X:", i)

			for _, b := range bytes[i:min(i+line, len(bytes))] {
				out += fmt.Sprintf(" %02X", b)
			}

			out += fmt.Sprintf("%*s", 3*max(i+line-len(bytes), 0), "")

			s := strings.Map(func(r rune) rune {
				if unicode.IsGraphic(r) {
					return r
				}
				return '.'
			}, string(bytes[i:min(i+line, len(bytes))]))

			if verbose {
				out += fmt.Sprintf("\t|%s|", s)
			}

			out += "\n"

			if !yield(out) {
				break
			}
		}
	}
}
