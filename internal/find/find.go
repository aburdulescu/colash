package find

import (
	"flag"
	"fmt"
	"os"
)

const usage = `Usage: find [-HL] [PATH]... [OPTIONS] [ACTIONS]

Search for files and perform actions on them.
First failed action stops processing of current file.

Defaults: PATH is current directory, action is '-print'

        -L,-follow      Follow symlinks
        -H              ...on command line only
        -xdev           Don't descend directories on other filesystems
        -maxdepth N     Descend at most N levels. -maxdepth 0 applies
                        actions to command line arguments only
        -mindepth N     Don't act on first N levels
        -depth          Act on directory *after* traversing it

Actions:
        ( ACTIONS )     Group actions for -o / -a
        ! ACT           Invert ACT's success/failure
        ACT1 [-a] ACT2  If ACT1 fails, stop, else do ACT2
        ACT1 -o ACT2    If ACT1 succeeds, stop, else do ACT2
                        Note: -a has higher priority than -o
        -name PATTERN   Match file name (w/o directory name) to PATTERN
        -iname PATTERN  Case insensitive -name
        -path PATTERN   Match path to PATTERN
        -ipath PATTERN  Case insensitive -path
        -regex PATTERN  Match path to regex PATTERN
        -type X         File type is X (one of: f,d,l,b,c,s,p)
        -executable     File is executable
        -perm MASK      At least one mask bit (+MASK), all bits (-MASK),
                        or exactly MASK bits are set in file's mode
        -mtime DAYS     mtime is greater than (+N), less than (-N),
                        or exactly N days in the past
        -mmin MINS      mtime is greater than (+N), less than (-N),
                        or exactly N minutes in the past
        -newer FILE     mtime is more recent than FILE's
        -inum N         File has inode number N
        -user NAME/ID   File is owned by given user
        -group NAME/ID  File is owned by given group
        -size N[bck]    File size is N (c:bytes,k:kbytes,b:512 bytes(def.))
                        +/-N: file size is bigger/smaller than N
        -links N        Number of links is greater than (+N), less than (-N),
                        or exactly N
        -prune          If current file is directory, don't descend into it

If none of the following actions is specified, -print is assumed
        -print          Print file name
        -print0         Print file name, NUL terminated
        -exec CMD ARG ; Run CMD with all instances of {} replaced by
                        file name. Fails if CMD exits with nonzero
        -exec CMD ARG + Run CMD with {} replaced by list of file names
        -quit           Exit
`

func Run(args []string) error {
	fset := flag.NewFlagSet("find", flag.ContinueOnError)
	fset.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}
	if err := fset.Parse(args); err != nil {
		return err
	}
	return nil
}
