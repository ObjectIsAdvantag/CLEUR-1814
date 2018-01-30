package main

import (
	"os"
	"time"

	"github.com/ObjectIsAdvantag/CLEUR-1814/4-cli/commands"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-spark"
	app.Version = "1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Patrick Riel",
			Email: "priel@cisco.com",
		},
	}
	app.Usage = "Example CLI For Cisco Spark"
	app.Commands = []cli.Command{
		commands.New(),
		commands.Delete(),
	}
	app.Run(os.Args)
}
