package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func runDirname(args []string) error {
	fset := flag.NewFlagSet("dirname", flag.ContinueOnError)
	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: dirname PATH...

Returns all but the last element of path, typically the path's directory.
`)
		os.Exit(1)
	}
	fset.Parse(args)
	if fset.NArg() < 1 {
		fset.Usage()
		return nil
	}
	for _, d := range fset.Args() {
		fmt.Println(filepath.Dir(d))
	}
	return nil
}
