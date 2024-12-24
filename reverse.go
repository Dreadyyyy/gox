package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func parseUnix(s string) (offs int64, bytes []byte, err error) {
	tokens := strings.Split(s, " ")

	_, err = fmt.Sscanf(strings.TrimSuffix(tokens[0], ":"), "%X", &offs)
	if err != nil {
		return
	}

	for _, t := range tokens[1:] {
		if len(t) != 2 {
			err = errors.New("Byte should be two a digit hexadecimal number")
			return
		}

		var b byte
		_, err = fmt.Sscanf(t, "%X", &b)
		if err != nil {
			return
		}

		bytes = append(bytes, b)
	}
	return
}

func reverse(r io.Reader, out *os.File) {
	br := bufio.NewReader(r)

	for ln, isPrefix, err := br.ReadLine(); err != io.EOF; ln, isPrefix, err = br.ReadLine() {
		if isPrefix {
			err = errors.New("Line is to long")
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		offs, bytes, err := parseUnix(string(ln))

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		if _, err := out.Seek(offs, 0); err != nil && offs < 0 {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		if _, err := out.Write(bytes); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			fmt.Println("what")
			os.Exit(1)
		}
	}
}
