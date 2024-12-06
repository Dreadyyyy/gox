package main

import (
	"fmt"
	"strings"
	"unicode"
)

func dumpSeq(bytes []byte) func(func(string) bool) {
	return func(yield func(string) bool) {
		for i := 0; i < len(bytes); i += 8 {
			out := fmt.Sprintf("%016X:", i)

			for _, b := range bytes[i:min(i+8, len(bytes))] {
				out += fmt.Sprintf(" %02X", b)
			}

			out += fmt.Sprintf("%*s", 3*max(i+8-len(bytes), 0), "")

			s := strings.Map(func(r rune) rune {
				if unicode.IsGraphic(r) {
					return r
				}
				return '.'
			}, string(bytes[i:min(i+8, len(bytes))]))
			out += fmt.Sprintf("\t|%s\n", s)

			if !yield(out) {
				break
			}
		}
	}
}
