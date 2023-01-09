package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

func runMkdir(args []string) error {
	fset := flag.NewFlagSet("mkdir", flag.ContinueOnError)

	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: mkdir [options] DIRECTORY...

Create given directory(ies).

Options:
`)
		fset.PrintDefaults()
		os.Exit(1)
	}

	p := fset.Bool("p", false, "no error if exists; make parent directories as needed")
	m := fset.String("m", "0755", "mode")

	if err := fset.Parse(args); err != nil {
		return err
	}

	if fset.NArg() < 1 {
		fset.Usage()
		return nil
	}
	mode, err := strconv.ParseUint(*m, 8, 32)
	if err != nil {
		return err
	}
	var maker func(string, fs.FileMode) error
	if *p {
		maker = os.MkdirAll
	} else {
		maker = os.Mkdir
	}
	for _, d := range fset.Args() {
		if err := maker(d, fs.FileMode(mode)); err != nil {
			return err
		}
	}
	return nil
}
