package main

import (
	"fmt"
	"strings"
)

func dump(s []byte) (out string) {
	for i := 0; i < len(s); i += 8 {
		out += fmt.Sprintf("%08x:", i)
		for _, b := range s[i:min(i+8, len(s))] {
			out += fmt.Sprintf(" %2x", b)
		}
		out += "\n"
	}
	return strings.ToUpper(out)
}
