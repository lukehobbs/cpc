package main

import (
	"fmt"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "commitArgs"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "Luke Hobbs",
			Email: "lukeehobbs@gmail.com",
		},
	}
	app.Usage = "Control CI pipeline using commit messages."
	app.UsageText = "commit -m \"<your commit message> commitArgs [options] [arguments]\""
	app.Commands = []cli.Command{}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
