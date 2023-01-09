package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func runSleep(args []string) error {
	fset := flag.NewFlagSet("sleep", flag.ContinueOnError)

	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: sleep NUMBER[SUFFIX]

Pause for NUMBER seconds.

SUFFIX may be 's' for seconds (the default), 'm' for minutes, 'h' for hours or 'd' for days.
NUMBER need not be an integer.
`)
		os.Exit(1)
	}

	if err := fset.Parse(args); err != nil {
		return err
	}

	if fset.NArg() < 1 {
		fset.Usage()
	}

	input := fset.Arg(0)
	var conv func(n float64) time.Duration
	switch {
	case strings.HasSuffix(input, "s"):
		input = input[:len(input)-1]
		conv = func(n float64) time.Duration {
			i := time.Duration(n)
			r := n - float64(i)
			return i*time.Second + time.Duration(r*1000)*time.Millisecond
		}
	case strings.HasSuffix(input, "m"):
		input = input[:len(input)-1]
		conv = func(n float64) time.Duration {
			i := time.Duration(n)
			r := n - float64(i)
			return i*time.Minute + time.Duration(r*60)*time.Second
		}
	case strings.HasSuffix(input, "h"):
		input = input[:len(input)-1]
		conv = func(n float64) time.Duration {
			i := time.Duration(n)
			r := n - float64(i)
			return i*time.Hour + time.Duration(r*60)*time.Minute
		}
	case strings.HasSuffix(input, "d"):
		input = input[:len(input)-1]
		conv = func(n float64) time.Duration {
			i := time.Duration(n)
			r := n - float64(i)
			return i*time.Hour*24 + time.Duration(r*24)*time.Hour
		}
	default:
		conv = func(n float64) time.Duration {
			i := time.Duration(n)
			r := n - float64(i)
			return i*time.Second + time.Duration(r*1000)*time.Millisecond
		}
	}
	n, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return err
	}
	time.Sleep(conv(n))
	return nil
}
