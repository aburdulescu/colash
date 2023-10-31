package main

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func runZip(args []string) error {
	fset := flag.NewFlagSet("zip", flag.ContinueOnError)

	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: zip [options] FILE...|DIRECTORY...

Zip given FILEs and/or DIRECTORYs.
Directories will be walked recursively.

Options:
`)
		fset.PrintDefaults()
		os.Exit(1)
	}

	fOutput := fset.String("o", "", "Output name")

	if err := fset.Parse(args); err != nil {
		return err
	}

	args = fset.Args()
	if len(args) == 0 {
		fset.Usage()
	}

	dstPath := *fOutput
	if dstPath == "" {
		if len(args) > 1 {
			return errors.New("need output name(-o) when specifying multiple inputs")
		}
		dstPath = args[0] + ".zip"
	}

	if filepath.Ext(dstPath) != ".zip" {
		dstPath += ".zip"
	}

	var files []string
	for _, arg := range args {
		fi, err := os.Stat(arg)
		if err != nil {
			return err
		}
		if fi.IsDir() {
			err := filepath.WalkDir(arg, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if d.IsDir() {
					return nil
				}
				files = append(files, path)
				return nil
			})
			if err != nil {
				return err
			}
		} else {
			files = append(files, filepath.Clean(arg))
		}
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	w := zip.NewWriter(dst)

	for _, file := range files {
		fmt.Println("adding:", file)

		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		if filepath.IsAbs(file) {
			file, _ = filepath.Rel("/", file)
		}

		fw, err := w.Create(file)
		if err != nil {
			return err
		}

		if _, err := io.Copy(fw, f); err != nil {
			return err
		}
	}

	return w.Close()
}
