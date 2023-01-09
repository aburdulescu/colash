package main

import (
	"fmt"
	"io"
	"io/fs"
)

func lsLongList(w io.Writer, fi fs.FileInfo, human bool) error {
	fmt.Fprintf(w, "%s", fi.Mode())
	if human {
		fmt.Fprintf(w, "\t%s", PrettySize(fi.Size()))
	} else {
		fmt.Fprintf(w, "\t%d", fi.Size())
	}
	fmt.Fprintf(w, "\t%s\t%d\t%02d:%02d\t%s\n",
		fi.ModTime().Month().String()[:3],
		fi.ModTime().Day(),
		fi.ModTime().Hour(),
		fi.ModTime().Minute(),
		fi.Name(),
	)
	return nil
}
