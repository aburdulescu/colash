package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func runRm(args []string) error {
	fset := flag.NewFlagSet("rm", flag.ContinueOnError)
	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: rm [OPTION] PATH...

Remove(unlink) PATHs.

Options:
`)
		fset.PrintDefaults()
		os.Exit(1)
	}
	r := fset.Bool("r", false, "remove directories and their contents recursively")
	f := fset.Bool("f", false, "ignore nonexistent files and arguments, never prompt")
	fset.Parse(args)
	for _, v := range fset.Args() {
		if !*f {
			answer := prompt("rm: remove '" + v + "'? [y/n]")
			switch answer {
			case "y":
			case "Y":
			case "yes":
			default:
				continue
			}
		}
		var err error
		if *r {
			err = os.RemoveAll(v)
		} else {
			err = os.Remove(v)
		}
		if *f {
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func prompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	fmt.Print(label + " ")
	s, _ = r.ReadString('\n')
	return strings.TrimSpace(s)
}
