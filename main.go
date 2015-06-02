package main

import (
	"github.com/brettchalupa/far/cmd"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "far"
	app.Usage = "find and replace text in files"
	app.Authors = []cli.Author{cli.Author{Name: "Brett Chalupa", Email: "brett@brettchalupa.com"}}
	app.Action = func(c *cli.Context) {
		if cmd.CheckArgs(c.Args()) {
			cmd.Execute(c.Args())
		}
	}

	app.Run(os.Args)
}
