package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func runId(args []string) error {
	fset := flag.NewFlagSet("id", flag.ContinueOnError)
	fset.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: id [OPTION]... USER

Print user and group information specified USER,\nor (when USER omitted) for the current user.

Options:
`)
		fset.PrintDefaults()
		os.Exit(1)
	}
	flagUser := fset.Bool("u", false, "print user ID")
	flagGroup := fset.Bool("g", false, "print group ID")
	flagSupGroups := fset.Bool("G", false, "print suplementary group IDs")
	flagPrintNames := fset.Bool("n", false, "print names instead of numbers")
	fset.Parse(args)
	var u *user.User
	if fset.NArg() < 1 {
		v, err := user.Current()
		if err != nil {
			return err
		}
		u = v
	} else {
		id := fset.Arg(0)
		var err error
		if _, err = strconv.ParseInt(id, 10, 32); err == nil {
			u, err = user.LookupId(id)
		} else {
			u, err = user.Lookup(id)
		}
		if err != nil {
			return err
		}
	}
	if *flagPrintNames &&
		((*flagUser && *flagGroup) || (*flagUser && *flagSupGroups) || (*flagGroup && *flagSupGroups)) {
		return errors.New("cannot print names for more than one choice")
	}
	if *flagUser {
		if *flagPrintNames {
			fmt.Printf("%s\n", u.Username)
		} else {
			fmt.Printf("%s\n", u.Uid)
		}
		return nil
	}
	g, err := user.LookupGroupId(u.Gid)
	if err != nil {
		return err
	}
	if *flagGroup {
		if *flagPrintNames {
			fmt.Printf("%s\n", g.Name)
		} else {
			fmt.Printf("%s\n", g.Gid)
		}
		return nil
	}
	groupsIds, err := u.GroupIds()
	if err != nil {
		return err
	}
	if *flagSupGroups {
		if *flagPrintNames {
			for i, id := range groupsIds {
				v, _ := user.LookupGroupId(id)
				fmt.Printf("%s", v.Name)
				if i != len(groupsIds)-1 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		} else {
			for i, id := range groupsIds {
				v, _ := user.LookupGroupId(id)
				fmt.Printf("%s", v.Gid)
				if i != len(groupsIds)-1 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		return nil
	}
	fmt.Printf("uid=%s(%s) gid=%s(%s) groups=", u.Uid, u.Username, u.Gid, g.Name)
	for i, id := range groupsIds {
		v, _ := user.LookupGroupId(id)
		fmt.Printf("%s(%s)", id, v.Name)
		if i != len(groupsIds)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println()
	return nil
}
