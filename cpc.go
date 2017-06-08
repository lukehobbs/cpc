package main

import (
	"fmt"
	"os"
	"sort"
	"time"
	"io/ioutil"
	"text/tabwriter"

	"github.com/urfave/cli"
	"github.com/ghodss/yaml"
	//"github.com/davecgh/go-spew/spew"
)

type AppFlags struct {
	BoolFlags 	[]cli.BoolFlag
	StringFlags []cli.StringFlag
	IntFlags 		[]cli.IntFlag
}

var Flags *AppFlags

func init() {
	// Find YAML file in project
	yamlFile, err := ioutil.ReadFile("cpc.yaml")
	if err != nil {
		fmt.Println("ERR: ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Flags)
	if err != nil {
		fmt.Println("ERR: ", err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "cpc"
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
	app.UsageText = "commit -m \"<your commit message> cpc [options] [arguments]\""
	app.Action = func(c *cli.Context) error {
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 0, '\t', 0)
		for _, b := range Flags.BoolFlags {
			fmt.Fprintf(w, "%s:\t%v\n", b.Name, c.IsSet(b.Name))
		}
		for _, s := range Flags.StringFlags {
			fmt.Fprintf(w, "%s:\t%s\n", s.Name, c.String(s.Name))
		}
		for _, i := range Flags.IntFlags {
			fmt.Fprintf(w, "%s:\t%d\n", i.Name, c.Int(i.Name))
		}
		w.Flush()
		// TODO: Create yaml/json file containing these variables for the pipeline to reference

		return nil
	}
	app.Flags = []cli.Flag{
		// Flags specific to cpc will go here
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

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	for i, s := range os.Args {
		if s == "cpc" {
			os.Args = os.Args[i:]
		}
	}
	app.Run(os.Args)
}
