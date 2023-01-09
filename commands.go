package main

import (
	"fmt"
	"os"
)

type Command struct {
	run  func([]string) error
	name string
}

var commands = []Command{
	{name: "pwd", run: runPwd},
	{name: "mkdir", run: runMkdir},
	{name: "rm", run: runRm},
	{name: "ls", run: runLs},
	{name: "cat", run: runCat},
	{name: "echo", run: runEcho},
	{name: "dirname", run: runDirname},
	{name: "basename", run: runBasename},
	{name: "id", run: runId},
	{name: "sleep", run: runSleep},
	{name: "zip", run: runZip},
	{name: "unzip", run: runUnzip},
	{name: "false", run: func([]string) error { os.Exit(1); return nil }},
	{name: "true", run: func([]string) error { os.Exit(0); return nil }},
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
