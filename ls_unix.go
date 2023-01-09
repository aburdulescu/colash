//go:build !windows

package main

import (
	"fmt"
	"io"
	"io/fs"
	"os/user"
	"strconv"
	"syscall"
)

func lsLongList(w io.Writer, fi fs.FileInfo) error {
	st := fi.Sys().(*syscall.Stat_t)
	u, err := user.LookupId(strconv.FormatUint(uint64(st.Uid), 10))
	if err != nil {
		return err
	}
	g, err := user.LookupGroupId(strconv.FormatUint(uint64(st.Gid), 10))
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\t%d\t%02d:%02d\t%s\n",
		fi.Mode(),
		u.Username,
		g.Name,
		fi.Size(),
		fi.ModTime().Month().String()[:3],
		fi.ModTime().Day(),
		fi.ModTime().Hour(),
		fi.ModTime().Minute(),
		fi.Name(),
	)
	return nil
}
