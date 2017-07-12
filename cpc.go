package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
// "text/tabwriter"
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
		// printFlags(c)
		fmt.Println(formatEnvs(c))
		return nil
	}
	app.Flags = []cli.Flag{
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
	// Set flag values from commit message.
	for i, s := range os.Args {
		if s == "cpc" {
			os.Args = os.Args[i:]
			return;
		}
		if strings.Contains(s, " cpc ") {
			a := strings.Split(os.Args[i], " ")
			for j, t := range a {
				if t == "cpc" {
					os.Args = a[j:]
				}
			}
			return;
		}
	}
	// log.Println("INFO: Commit message does not contain 'cpc'.")
	return;
}

func formatEnvs(c *cli.Context) string {
	var b bytes.Buffer
	var a []string
	for _, f := range flags.BoolFlags {
		s := splitName(f.Name)
		v := c.IsSet(s[0])
		a = append(a, fmt.Sprintf("%s=%s", strings.ToUpper(s[0]), strconv.FormatBool(v)))
	}
	for _, f := range flags.StringFlags {
		s := splitName(f.Name)
		v := c.String(s[0])
			a = append(a, fmt.Sprintf("%s=%s", strings.ToUpper(s[0]), v))
	}
	for _, f := range flags.IntFlags {
		s := splitName(f.Name)
		v := c.Int(s[0])
		a = append(a, fmt.Sprintf("%s=%s", strings.ToUpper(s[0]), strconv.Itoa(v)))
	}
	b.WriteString(strings.Join(a, ", "))
	return b.String()
}

func splitName(s string) []string {
	return strings.Split(s, ",")
}
