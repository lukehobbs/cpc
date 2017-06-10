package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/ghodss/yaml"
	"github.com/urfave/cli"
)

type appFlags struct {
	BoolFlags   []cli.BoolFlag
	StringFlags []cli.StringFlag
	IntFlags    []cli.IntFlag
}

var flags *appFlags

func init() {
	// Check if yaml file exists in root of project.
	if _, err := os.Stat("cpc.yaml"); os.IsNotExist(err) {
		log.Fatal("ERROR: cpc.yaml must exist in the root directory of your project.")
	}
	yamlFile, err := ioutil.ReadFile("cpc.yaml")
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	// Unmarshal cpc.yaml into $yamlFile.
	err = yaml.Unmarshal(yamlFile, &flags)
	if err != nil {
		log.Fatal("ERROR: ", err)
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
		printFlags(c)
		exportEnvs(c)
		return nil
	}
	app.Flags = []cli.Flag{
	// Flags specific to cpc will go here.
	// i.e. --json, --yaml, --envs for output format.
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	setFlags(app)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func setFlags(app *cli.App) {
	// Set app.Flags from yaml file
	for _, b := range flags.BoolFlags {
		app.Flags = append(app.Flags, b)
	}
	for _, s := range flags.StringFlags {
		app.Flags = append(app.Flags, s)
	}
	for _, i := range flags.IntFlags {
		app.Flags = append(app.Flags, i)
	}
	// Check if commit message contains 'cpc'.
	if !contains(os.Args, "cpc") {
		log.Println("WARN: Commit message does not contain 'cpc'.")
		os.Exit(1)
	}
	// Set flag values from commit message.
	for i, s := range os.Args {
		if s == "cpc" {
			os.Args = os.Args[i:]
		}
	}
}

func printFlags(c *cli.Context) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	for _, f := range flags.BoolFlags {
		s := splitName(f.Name)
		v := c.IsSet(s[0])
		_, err := fmt.Fprintf(w, "%s:\t\t\t  %v\n", s[0], v)
		if err != nil {
			log.Fatal("ERROR: ", err)
		}
	}
	for _, f := range flags.StringFlags {
		s := splitName(f.Name)
		v := c.String(s[0])
		_, err := fmt.Fprintf(w, "%s:\t\t\t  %s\n", s[0], v)
		if err != nil {
			log.Fatal("ERROR: ", err)
		}
	}
	for _, f := range flags.IntFlags {
		s := splitName(f.Name)
		v := c.Int(s[0])
		_, err := fmt.Fprintf(w, "%s:\t\t\t  %v\n", s[0], v)
		if err != nil {
			log.Fatal("ERROR: ", err)
		}
	}
	err := w.Flush()
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func exportEnvs(c *cli.Context) {
	fo, err := os.OpenFile("cpc-out.sh", os.O_RDWR|os.O_CREATE, 0744)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	_, err = fmt.Fprintf(fo, "#! /usr/bin/env bash\n")
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	for _, f := range flags.BoolFlags {
		s := splitName(f.Name)
		v := c.IsSet(s[0])
		_, err := fmt.Fprintf(fo, "export %s=\"%s\"\n", strings.ToUpper(s[0]), strconv.FormatBool(v))
		if err != nil {
			log.Fatal("ERROR: ", err)
		}
	}
	for _, f := range flags.StringFlags {
		s := splitName(f.Name)
		v := c.String(s[0])
		_, err := fmt.Fprintf(fo, "export %s=\"%s\"\n", strings.ToUpper(s[0]), v)
		if err != nil {
			log.Fatal("ERROR: ", err)
		}
	}
	for _, f := range flags.IntFlags {
		s := splitName(f.Name)
		v := c.Int(s[0])
		_, err := fmt.Fprintf(fo, "export %s=\"%s\"\n", strings.ToUpper(s[0]), strconv.Itoa(v))
		if err != nil {
			log.Fatal("ERROR: ", err)
		}
	}
	if err := fo.Close(); err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func splitName(s string) []string {
	return strings.Split(s, ",")
}

func contains(s []string, b string) bool {
	for _, a := range s {
		if a == b {
			return true
		}
	}
	return false
}
