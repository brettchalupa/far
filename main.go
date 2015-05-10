package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "far"
	app.Usage = "find and replace text in files"
	app.Authors = []cli.Author{cli.Author{Name: "Brett Chalupa", Email: "brett@brettchalupa.com"}}
	app.Action = func(c *cli.Context) {
		if checkArgs(c.Args()) {
			if fileExists(c.Args()[0]) {
				count, err := findAndReplace(c.Args()[0], c.Args()[1], c.Args()[2])
				if err != nil {
					panic(err)
				}
				fmt.Println("Times replaced", count)
			} else {
				fmt.Println("Provided file does not exist")
			}
		} else {
			return
		}
	}

	app.Run(os.Args)
}

func checkArgs(args cli.Args) bool {
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

func fileExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	} else {
		return true
	}
}

func findAndReplace(path, current, update string) (int, error) {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	write_file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer write_file.Close()

	writer := bufio.NewWriter(file)

	replaceCount := 0

	for _, line := range lines {
		if strings.Contains(line, current) {
			line = strings.Replace(line, current, update, -1)
			replaceCount++
		}
		write_file.WriteString(line + "\n")
	}

	writer.Flush()

	return replaceCount, nil
}
