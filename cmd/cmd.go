package cmd

import (
	"fmt"
	"github.com/brettchalupa/far/far"
	"github.com/codegangsta/cli"
)

func CheckArgs(args cli.Args) bool {
	var neededArgs bool

	argsCount := len(args)

	if argsCount >= 3 {
		neededArgs = true
	} else {
		if argsCount == 0 {
			fmt.Println("Please provide the path, original value, and new value")
		} else if argsCount == 1 {
			fmt.Println("Please provide the original value and new value")
		} else if argsCount == 2 {
			fmt.Println("Please provide the new value")
		}

		neededArgs = false
	}

	return neededArgs
}

func Execute(args cli.Args) {
	if far.FileExists(args[0]) {
		count, err := far.FindAndReplace(args[0], args[1], args[2])
		if err != nil {
			panic(err)
		}
		fmt.Println("Times replaced", count)
	} else {
		fmt.Println("Provided file does not exist")
	}
}
