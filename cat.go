package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func runCat(args []string) error {
	fset := flag.NewFlagSet("cat", flag.ContinueOnError)

	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: cat [FILE]...

Print FILEs to stdout.

With no FILE, or when FILE is -, read standard input.
`)
		os.Exit(1)
	}

	if err := fset.Parse(args); err != nil {
		return err
	}

	var r io.Reader
	if fset.NArg() == 0 {
		r = os.Stdin
	} else {
		if fset.Arg(0) == "-" {
			r = os.Stdin
		} else {
			f, err := os.Open(fset.Arg(0))
			if err != nil {
				return err
			}
			defer f.Close()
			r = f
		}
	}
	_, err := io.Copy(os.Stdout, r)
	return err
}
