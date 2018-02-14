package main

import (
	"os"
	"github.com/urfave/cli"
	"fmt"
	"time"
	"github.com/yoichiro/dialogflow-query-checker/config"
	"github.com/yoichiro/dialogflow-query-checker/check"
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
			ArgsUsage: "[configuration file] A configuration YAML file which has conditions and expected query results",
			Action: func(c *cli.Context) error {
				start := time.Now()
				if !c.Args().Present() {
					return cli.NewExitError("A configuration file not specified", 1)
				}
				path := c.Args().First()
				def, err := config.LoadConfigurationFile(path)
				if err != nil {
					return cli.NewExitError(err, 1)
				}
				fmt.Printf("The configuration file loaded: %s\n", path)
				fmt.Println("Running query tests for Dialogflow.")
				results, err := check.Execute(def)
				fmt.Println()
				if err != nil {
					return cli.NewExitError(err, 1)
				}
				end := time.Now()
				if results.Len() == 0 {
					fmt.Printf("Finished in %f seconds.\n", (end.Sub(start)).Seconds())
					fmt.Println("All tests passed.")
					return nil
				} else {
					for e := results.Front(); e != nil; e = e.Next() {
						fmt.Printf("%s\n", e.Value)
					}
					fmt.Printf("Finished in %f seconds.\n", (end.Sub(start)).Seconds())
					return cli.NewExitError(fmt.Sprintf("%d tests failed.", results.Len()), 1)
				}
			},
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command [%q] not found.\n", command)
	}
	app.Run(os.Args)
}
