package main

import (
	"flag"
	"fmt"
	"os"
)

func runPwd(args []string) error {
	fset := flag.NewFlagSet("pwd", flag.ContinueOnError)
	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: pwd

Print the full path of the current working directory
`)
		os.Exit(1)
	}
	fset.Parse(args)
	p, err := os.Getwd()
	if err != nil {
		die(err)
	}
	fmt.Println(p)
	return nil
}
