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

func reverseUnix(in io.Reader, out *os.File) {
	br := bufio.NewReader(in)

	_, err := out.Seek(0, 0)
	seekable := err == nil

	for ln, isPrefix, err := br.ReadLine(); err != io.EOF; ln, isPrefix, err = br.ReadLine() {
		if isPrefix {
			err = errors.New("Line is to long")
		}

		if err != nil {
			fatal(err.Error())
		}

		offs, bytes, err := parseUnix(string(ln))

		if err != nil {
			fatal(err.Error())
		}

		if seekable {
			if _, err := out.Seek(offs, 0); err != nil {
				fatal(err.Error())
			}
		}

		if _, err := out.Write(bytes); err != nil {
			fatal(err.Error())
		}
	}
}

func reversePlain(in io.Reader, out *os.File) {
	buff := make([]byte, 1)

	curr := ""
	for n, err := in.Read(buff); n != 0; n, err = in.Read(buff) {
		if err != nil {
			fatal(err.Error())
		}

		if buff[0] == '\n' || buff[0] == ' ' {
			continue
		}

		curr += string(buff[0])

		if len(curr) == 2 {
			var b byte
			if _, err := fmt.Sscanf(curr, "%X", &b); err != nil {
				fatal(err.Error())
			}

			if _, err := out.Write([]byte{b}); err != nil {
				fatal(err.Error())
			}

			curr = ""
		}
	}
}

func reverse(in io.Reader, out *os.File, plain bool) {
	if plain {
		reversePlain(in, out)
	} else {
		reverseUnix(in, out)
	}
}
