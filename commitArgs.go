package main

import (
  "fmt"
	"os"
	"sort"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "commitArgs"
	app.Version = "0.0.1"
	app.HideHelp = true
	app.HideVersion = true
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "Luke Hobbs",
			Email: "lukeehobbs@gmail.com",
		},
	}
	app.Usage = "Control CI pipeline using commit messages."
	app.UsageText = "commit -m \"<your commit message> commitArgs [options] [arguments]\""
	app.Action = func(c *cli.Context) error {
    fmt.Printf("Run full pipeline?\t%-5v\n", c.IsSet("full-pipeline"))
    fmt.Printf("Leave stack running?\t%-5v\n", c.IsSet("leave-up"))
    fmt.Printf("Run Serverspec tests?\t%-5v\n", c.IsSet("serverspec"))
    return nil
  }
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "full-pipeline"},
    cli.BoolFlag{Name: "leave-up"},
    cli.BoolFlag{Name: "serverspec"},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	for i, s := range os.Args {
		if s == "commitArgs" {
			os.Args = os.Args[i:]
		}
	}
	app.Run(os.Args)
}
