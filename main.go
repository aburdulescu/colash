package main

import (
	"flag"
	"fmt"
	"os"
)

const progname = "colash"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s [options] command [args]

Options:
`, progname)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
For help on individual command, run '%s command -h'.
`, progname)
		os.Exit(1)
	}
	listCommands := flag.Bool("l", false, "list available commands")
	flag.Parse()
	if *listCommands {
		printCommands()
		return
	}
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
	}
	run := findCommand(args[0])
	if run == nil {
		die("unknown command")
	}
	if err := run(args[1:]); err != nil {
		die(err)
	}
}

func die(a ...any) {
	fmt.Fprintln(os.Stderr, "error:", fmt.Sprint(a...))
	os.Exit(1)
}
