package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	USAGE = `Usage:
    gox [options] [infile[outfile]]
Options:
    -h | -help  Show help
    -v |        Verbose
    -c |        Format <cols> bytes per line
    `
	COLS = "Invalid number of columns: max 256\n"
)

func main() {
	var err error

	h := false
	flag.BoolVar(&h, "h", false, "Show help")
	flag.BoolVar(&h, "help", false, "Show help")
	v := flag.Bool("v", false, "Verbose")
	c := flag.Int("c", 8, "Format <cols> bytes per line")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, USAGE)
	}

	flag.Parse()

	if *c > 256 {
		fmt.Fprint(os.Stderr, COLS)
		os.Exit(1)
	}

	args := flag.Args()

	if len(args) > 2 || h {
		fmt.Fprint(os.Stderr, USAGE)
		os.Exit(1)
	}

	in, out := os.Stdin, os.Stdout

	if len(args) > 0 {
		in, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
	}

	if len(args) == 2 {
		out, err = os.Create(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
	}

	for l, err := range unixSeq(in, *v, *c) {
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}

		fmt.Fprint(out, l)
	}
}
