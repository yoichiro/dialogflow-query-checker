package main

import (
	"os"
	"github.com/urfave/cli"
	"fmt"
	"time"
	"github.com/yoichiro/dialogflow-query-checker/config"
	"github.com/yoichiro/dialogflow-query-checker/check"
	"github.com/yoichiro/dialogflow-query-checker/output"
)

func main() {
	app := cli.NewApp()
	app.Name = "dialogflow-query-checker"
	app.Usage = "A query check tool for Dialogflow"
	app.UsageText = "dialogflow-query-checker [global options] command [arguments]"
	app.Version = VERSION
	app.Commands = []cli.Command{
		{
			Name: "run",
			Usage: "Execute posting each query and checking it based on a configuration file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "output, o",
					Usage: "Output results as JUnit XML format `FILE`",
				},
				cli.BoolFlag{
					Name: "debug, d",
					Usage: "Debug mode",
				},
				cli.IntFlag{
					Name:  "retry, r",
					Usage: "The count of retrying",
				},
			},
			ArgsUsage: "[configuration file] A configuration YAML file which has conditions and expected query results",
			Action: func(c *cli.Context) error {
				fmt.Printf("%s version %s\n", app.Name, app.Version)

				start := time.Now()

				if !c.Args().Present() {
					return cli.NewExitError("[Error] A configuration file not specified", 1)
				}

				path := c.Args().First()
				def, err := config.LoadConfigurationFile(path)
				if err != nil {
					return cli.NewExitError(fmt.Sprintf("[Error] %s", err), 1)
				}
				fmt.Printf("The configuration file loaded: %s\n", path)

				debug := c.Bool("debug")
				if debug {
					fmt.Println("The debug mode is on")
				}
				(&def.Environment).Debug = debug

				retryCount := c.Int("retry")
				if retryCount < 0 {
					return cli.NewExitError(fmt.Sprint("[Error] The count of retrying must be positive"), 1)
				}
				fmt.Printf("The count of retrying: %d\n", c.Int("retry"))
				(&def.Environment).RetryCount = retryCount

				fmt.Println("Running query tests for Dialogflow.")
				holder, err := check.Execute(def)
				fmt.Println()
				if err != nil {
					return cli.NewExitError(fmt.Sprintf("[Error] %s", err), 1)
				}

				end := time.Now()

				output.Standard(holder, start, end)
				if c.String("output") != "" {
					err = output.JunitXml(holder, c.String("output"), start, end)
					if err != nil {
						return cli.NewExitError(fmt.Sprintf("[Error] %s", err), 1)
					}
				}

				if holder.AllFailureAssertResultCount() == 0 {
					return nil
				} else {
					failureTestResultCount := holder.AllFailureTestResultCount()
					if failureTestResultCount == 1 {
						return cli.NewExitError(fmt.Sprint("1 test failed.", ), 1)
					} else {
						return cli.NewExitError(fmt.Sprintf("%d tests failed.", failureTestResultCount), 1)
					}
				}
			},
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command [%q] not found.\n", command)
	}
	app.Run(os.Args)
}
