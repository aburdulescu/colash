package main

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

func runUnzip(args []string) error {
	fset := flag.NewFlagSet("unzip", flag.ContinueOnError)

	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: uzip [options] FILE[.zip]

Unzip given FILE.

Options:
`)
		fset.PrintDefaults()
		os.Exit(1)
	}

	fDir := fset.String("d", "", "Directory where to extract the files.")
	fList := fset.Bool("l", false, "List archive files.")

	if err := fset.Parse(args); err != nil {
		return err
	}

	args = fset.Args()
	if len(args) == 0 {
		fset.Usage()
	}

	r, err := zip.OpenReader(args[0])
	if err != nil {
		return err
	}
	defer r.Close()

	if *fDir != "" {
		if err := os.MkdirAll(*fDir, 0755); err != nil {
			return err
		}
	}

	base := *fDir

	switch {
	case *fList:
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
		defer w.Flush()
		fmt.Fprintln(w, "Name\tSize")
		fmt.Fprintln(w, "----\t----")
		for _, f := range r.File {
			fh := f.FileHeader
			fmt.Fprintf(w, "%s\t%d\n", fh.Name, fh.UncompressedSize64)
		}
	default:
		for _, f := range r.File {
			fmt.Println("extracting:", f.Name)
			src, err := f.Open()
			if err != nil {
				return err
			}
			defer src.Close()
			dst, err := os.Create(filepath.Join(base, strings.ReplaceAll(f.Name, "..", "")))
			if err != nil {
				return err
			}
			defer dst.Close()

			// copy in chunks to avoid decompression bomb
			for {
				_, err := io.CopyN(dst, src, 8192)
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}
			}
		}
	}

	return nil
}
