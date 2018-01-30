package commands

import (
	"fmt"
	"os"

	"github.com/ObjectIsAdvantag/CLEUR-1814/4-cli/ciscosparkapi/teams"
	"github.com/urfave/cli"
)

func New() cli.Command {
	return cli.Command{
		Name:        "new",
		ShortName:   "n",
		Description: "performs post API calls to the ciscospark API",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "team", Usage: "spark team name"},
			cli.StringFlag{Name: "token", Usage: "spark access token"},
		},
		Action: func(c *cli.Context) error {
			token := c.String("token")
			if token == "" {
				token = os.Getenv("CISCO_SPARK_ACCESS_TOKEN")
			}
			res, err := teams.Create(c.String("team"), token)
			if err != nil {
				fmt.Fprintf(c.App.Writer, err.Error()+"\n")
				return err
			}
			msg := fmt.Sprintf("created: %v\n", res.Created)
			msg += fmt.Sprintf("name: %v\n", res.Name)
			msg += fmt.Sprintf("id: %v\n", res.ID)

			fmt.Fprintf(c.App.Writer, msg)
			return nil
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			fmt.Println(err.Error())
			fmt.Fprintf(c.App.Writer, "for shame\n")
			return err
		},
	}
}

func Delete() cli.Command {
	return cli.Command{
		Name:        "delete",
		ShortName:   "d",
		Description: "performs delete API calls to the ciscospark API",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "teamID", Usage: "spark team id"},
			cli.StringFlag{Name: "token", Usage: "spark access token"},
		},
		Action: func(c *cli.Context) error {

			token := c.String("token")
			if token == "" {
				token = os.Getenv("CISCO_SPARK_ACCESS_TOKEN")
			}

			err := teams.Delete(c.String("teamID"), token)
			if err != nil {
				fmt.Fprintf(c.App.Writer, err.Error()+"\n")
				return err
			}
			msg := fmt.Sprintf("success\n")
			fmt.Fprintf(c.App.Writer, msg)
			return nil
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			fmt.Println(err.Error())
			fmt.Fprintf(c.App.Writer, "for shame\n")
			return err
		},
	}
}
