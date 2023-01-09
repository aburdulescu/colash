package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func runBasename(args []string) error {
	fset := flag.NewFlagSet("basename", flag.ContinueOnError)
	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: dirname NAME [SUFFIX]\n   or: mo dirname OPTION... NAME...

Print NAME with any leading directory components removed.

If specified, also remove a trailing SUFFIX.

Options:
`)
		fset.PrintDefaults()
		fmt.Fprintln(os.Stderr, `
Examples:
  basename /usr/bin/sort          -> "sort"
  basename include/stdio.h .h     -> "stdio"
  basename -s .h include/stdio.h  -> "stdio"
  basename -a any/str1 any/str2   -> "str1" followed by "str2"`)
		os.Exit(1)
	}
	a := fset.Bool("a", false, "support multiple arguments and treat each as a NAME")
	s := fset.String("s", "", "remove a trailing SUFFIX; implies -a")
	fset.Parse(args)
	if fset.NArg() < 1 {
		fset.Usage()
		return nil
	}
	if *s != "" {
		for _, d := range fset.Args() {
			fmt.Println(strings.TrimSuffix(filepath.Base(d), *s))
		}
		return nil
	}
	if *a {
		for _, d := range fset.Args() {
			fmt.Println(filepath.Base(d))
		}
		return nil
	}
	if fset.NArg() == 2 {
		fmt.Println(strings.TrimSuffix(filepath.Base(fset.Arg(0)), fset.Arg(1)))
	} else {
		fmt.Println(filepath.Base(fset.Arg(0)))
	}
	return nil
}
