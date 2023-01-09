package main

import (
	"flag"
	"fmt"
	"os"
)

func runEcho(args []string) error {
	fset := flag.NewFlagSet("echo", flag.ContinueOnError)
	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: echo [OPTION] STRING...

Display a line of text.

Options:
`)
		fset.PrintDefaults()
		os.Exit(1)
	}
	n := fset.Bool("n", false, "do not output the trailing newline")
	// TODO: e := fset.Bool("e", false, "enable interpretation of backslash escapes")
	fset.Parse(args)
	for _, v := range fset.Args() {
		fmt.Print(v)
		if !*n {
			fmt.Println()
		}
	}
	return nil
}
