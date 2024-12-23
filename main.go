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
    -p |        Plain hex dump output style
    -v |        Verbose
    -c |        Format <cols> bytes per line
    -o |        Add <offset> to the displayed file position
    -r |        Reverse
    `
	COLS = "Invalid number of columns: max 256"
)

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {
	var err error

	h := flag.Bool("help", false, "Show help")
	p := flag.Bool("p", false, "Plain hex dump output style")
	v := flag.Bool("v", false, "Verbose")
	c := flag.Int("c", 8, "Format <cols> bytes per line")
	o := flag.Int("o", 0, "Add offset")
	r := flag.Bool("r", false, "Reverse")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, USAGE)
	}

	flag.Parse()

	if *c > 256 {
		fatal(COLS)
	}

	args := flag.Args()

	if len(args) > 2 || *h {
		fatal(USAGE)
	}

	in, out := os.Stdin, os.Stdout

	if len(args) > 0 {
		in, err = os.Open(args[0])
		if err != nil {
			fatal(err.Error())
		}
		defer in.Close()
	}

	if len(args) == 2 {
		out, err = os.OpenFile(args[1], os.O_WRONLY, 0644)
		if err != nil {
			fatal(err.Error())
		}
		defer out.Close()
	}

	if *r {
		reverse(in, out, *p)
	} else {
		dump(in, out, *p, *v, *c, *o)
	}
}
