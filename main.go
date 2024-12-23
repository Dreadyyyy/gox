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
    -ps|        Plain hex dump output style
    -v |        Verbose
    -c |        Format <cols> bytes per line
    -o |        Add <offset> to the displayed file position
    `
	COLS = "Invalid number of columns: max 256"
)

func main() {
	var err error

	h := flag.Bool("help", false, "Show help")
	ps := flag.Bool("ps", false, "Plain hex dump output style")
	v := flag.Bool("v", false, "Verbose")
	c := flag.Int("c", 8, "Format <cols> bytes per line")
	o := flag.Int("o", 0, "Add offset")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, USAGE)
	}

	flag.Parse()

	if *c > 256 {
		fmt.Fprintln(os.Stderr, COLS)
		os.Exit(1)
	}

	args := flag.Args()

	if len(args) > 2 || *h {
		fmt.Fprintln(os.Stderr, USAGE)
		os.Exit(1)
	}

	in, out := os.Stdin, os.Stdout

	if len(args) > 0 {
		in, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	if len(args) == 2 {
		out, err = os.Create(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	for l, err := range dumpSeq(in, *ps, *v, *c, *o) {
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Fprint(out, l)
	}
}
