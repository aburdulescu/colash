package main

import (
	"fmt"
	"os"
)

type Command struct {
	name string
	run  func([]string) error
}

var commands = []Command{
	{"pwd", runPwd},
	{"mkdir", runMkdir},
	{"rm", runRm},
	{"ls", runLs},
	{"cat", runCat},
	{"echo", runEcho},
	{"dirname", runDirname},
	{"basename", runBasename},
	{"id", runId},
	{"sleep", runSleep},
	{"zip", runZip},
	{"unzip", runUnzip},
	{"false", func([]string) error { os.Exit(1); return nil }},
	{"true", func([]string) error { os.Exit(0); return nil }},
}

func printCommands() {
	for _, cmd := range commands {
		fmt.Println(cmd.name)
	}
}

func findCommand(name string) func([]string) error {
	for _, cmd := range commands {
		if cmd.name == name {
			return cmd.run
		}
	}
	return nil
}
