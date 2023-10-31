package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

func runLs(args []string) error {
	fset := flag.NewFlagSet("ls", flag.ContinueOnError)

	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: ls [OPTION] PATH...

List directory contents.

Options:
`)
		fset.PrintDefaults()
		os.Exit(1)
	}

	one := fset.Bool("1", false, "list one file per line")
	a := fset.Bool("a", false, "include entries which start with .")
	// A := fset.Bool("A", false, "like -a, but exclude . and ..")
	// x := fset.Bool("x", false, "list by lines")
	// d := fset.Bool("d", false, "list directory entries instead of contents")
	// L := fset.Bool("L", false, "follow symlinks")
	// H := fset.Bool("H", false, "follow symlinks on command line")
	// R := fset.Bool("R", false, "recurse")
	// p := fset.Bool("p", false, "append / to dir entries")
	// F := fset.Bool("F", false, "append indicator (one of */=@|) to entries")
	l := fset.Bool("l", false, "long listing format")
	// i := fset.Bool("i", false, "list inode numbers")
	// s := fset.Bool("s", false, "list allocated blocks")
	// lc := fset.Bool("lc", false, "list ctime")
	// lu := fset.Bool("lu", false, "list atime")
	// fullTime := fset.Bool("full-time", false, "list full date and time")
	h := fset.Bool("h", false, "human readable sizes (1K 243M 2G)")
	S := fset.Bool("S", false, "sort by size")
	// X := fset.Bool("X", false, "sort by extension")
	// v := fset.Bool("v", false, "sort by version")
	t := fset.Bool("t", false, "sort by mtime")
	// tc := fset.Bool("tc", false, "sort by ctime")
	// tu := fset.Bool("tu", false, "sort by atime")
	r := fset.Bool("r", false, "reverse sort order")
	// w := fset.Uint("w", 0, "format N columns wide")
	// color := fset.String("color", "always", "control coloring, one of: always,never,auto")

	if err := fset.Parse(args); err != nil {
		return err
	}

	sep := "  "
	if *one {
		sep = "\n"
	}
	// TODO: work on multiple args
	dir := "."
	if fset.NArg() > 0 {
		dir = fset.Arg(0)
	}
	argIsFile := false
	var result []fs.FileInfo
	if fi, err := os.Stat(dir); err == nil {
		if fi.IsDir() {
			entries, err := os.ReadDir(dir)
			if err != nil {
				return err
			}
			result = make([]fs.FileInfo, 0, len(entries))
			for _, entry := range entries {
				info, err := entry.Info()
				if err != nil {
					return err
				}
				result = append(result, info)
			}
		} else {
			argIsFile = true
			result = append(result, fi)
		}
	} else {
		return err
	}
	if !argIsFile {
		if *S {
			sort.Slice(result, func(i, j int) bool {
				if *r {
					return result[i].Size() < result[j].Size()
				}
				return result[i].Size() > result[j].Size()
			})
		}
		if *t {
			sort.Slice(result, func(i, j int) bool {
				if *r {
					return !result[i].ModTime().After(result[j].ModTime())
				}
				return !result[i].ModTime().Before(result[j].ModTime())
			})
		}
	}
	if !argIsFile {
		if *a {
			if !*l {
				fmt.Print(".", sep, "..", sep)
			} else {
				fi1, _ := os.Stat(".")
				fi2, _ := os.Stat("..")
				var tmp []fs.FileInfo
				tmp = append(tmp, fi1, fi2)
				tmp = append(tmp, result...)
				result = tmp
			}
		}
	}
	if *l {
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', 0)
		defer w.Flush()
		for _, fi := range result {
			if strings.HasPrefix(fi.Name(), ".") {
				if !argIsFile && !*a {
					continue
				}
			}
			if err := lsLongList(w, fi, *h); err != nil {
				return err
			}
		}
		return nil
	}
	for _, fi := range result {
		if strings.HasPrefix(fi.Name(), ".") {
			if !argIsFile && !*a {
				continue
			}
		}
		fmt.Print(fi.Name(), sep)
	}
	if !*one {
		fmt.Println()
	}
	return nil
}
