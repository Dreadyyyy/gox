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

func dump(in io.Reader, out *os.File, p bool, v bool, cols int, offs int) {
	format := " %02X"
	padding := 3
	if p {
		format = "%02X"
		padding = 2
	}

	bytes := make([]byte, cols)

	for n, err := in.Read(bytes); n != 0; n, err = in.Read(bytes) {
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		var sb strings.Builder

		if !p {
			sb.WriteString(fmt.Sprintf("%016X", offs))
		}

		for _, b := range bytes[:n] {
			sb.WriteString(fmt.Sprintf(format, b))
		}
		sb.WriteString(fmt.Sprintf("%*s", padding*(cols-n), ""))

		if v {
			sb.WriteString("\t")
			verbose(&sb, bytes[:n])
		}

		sb.WriteString("\n")

		fmt.Fprint(out, sb.String())

		offs += n
	}
}
