package main

import (
	"fmt"
	"io"
	"io/fs"
)

func lsLongList(w io.Writer, fi fs.FileInfo) error {
	fmt.Fprintf(w, "%s\t%d\t%s\t%d\t%02d:%02d\t%s\n",
		fi.Mode(),
		fi.Size(),
		fi.ModTime().Month().String()[:3],
		fi.ModTime().Day(),
		fi.ModTime().Hour(),
		fi.ModTime().Minute(),
		fi.Name(),
	)
	return nil
}
