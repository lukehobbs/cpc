package main

import (
	"fmt"
	"os"
	"sort"
	"time"
	"io/ioutil"

	"github.com/urfave/cli"
	"github.com/ghodss/yaml"
	"github.com/davecgh/go-spew/spew"
)

type AppFlags struct {
	BoolFlags 	[]cli.BoolFlag
	StringFlags []cli.StringFlag
	IntFlags 		[]cli.IntFlag
}

var Flags *AppFlags

func init() {
	// Find YAML file in project
	yamlFile, err := ioutil.ReadFile("commitArgs.yaml")
	if err != nil {
		fmt.Println("ERR: ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Flags)
	if err != nil {
		fmt.Println("ERR: ", err)
	}
	// Unmarshal YAML into app.Flags
	spew.Dump(Flags)
}

func main() {
	app := cli.NewApp()
	app.Name = "commitArgs"
	app.Version = "0.1.0"
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
		// fmt.Printf("Run full pipeline:\t\t%v\n", c.IsSet("full-pipeline"))
		// fmt.Printf("Leave stack running:\t\t%v\n", c.IsSet("leave-up"))
		// fmt.Printf("Stack should stay up for:\t%d seconds\n", c.Int("run-time"))
		// fmt.Printf("Run Serverspec tests:\t\t%v\n", c.IsSet("serverspec"))
		// fmt.Printf("Using profile:\t\t\t%s\n", c.String("it-me"))
		// // TODO: Create yaml/json file containing these variables for the pipeline to reference
		// y, err := yaml.Marshal(app.Flags)
		// if err != nil {
		// 	fmt.Println("Err: ", err)
		// }
		// fmt.Println(string(y))
		return nil
	}
	for _, b := range Flags.BoolFlags {
		app.Flags = append(app.Flags, b)
	}
	for _, s := range Flags.StringFlags {
		app.Flags = append(app.Flags, s)
	}
	for _, i := range Flags.IntFlags {
		app.Flags = append(app.Flags, i)
	}
	// app.Flags = []cli.Flag{Flags.BoolFlags, Flags.StringFlags, Flags.IntFlags}
		// []cli.Flag{
		// cli.BoolFlag{Name: "full-pipeline"},
		// cli.BoolFlag{Name: "leave-up"},
		// cli.BoolFlag{Name: "serverspec"},
		// cli.StringFlag{Name: "it-me"},
		// cli.IntFlag{Name: "run-time"},
		// }
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	for i, s := range os.Args {
		if s == "commitArgs" {
			os.Args = os.Args[i:]
		}
	}
	app.Run(os.Args)
}
