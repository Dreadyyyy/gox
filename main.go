package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var err error

	if len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "Usage: gox [infile[outfile]]\n")
		os.Exit(1)
	}

	in, out := os.Stdin, os.Stdout

	if len(os.Args) >= 2 {
		in, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
	}

	if len(os.Args) == 3 {
		out, err = os.Create(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
	}

	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, in); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(out, dump(buf.Bytes()))
}
