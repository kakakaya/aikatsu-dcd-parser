package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/kakakaya/aikatsu-dcd-parser"
	"github.com/urfave/cli"
)

const (
	loginNotRequiredCategoryLabel string = "Login not required commands"
	loginRequiredCategoryLabel    string = "Login required commands"
)

func main() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}

	app := cli.NewApp()
	app.Name = "dcdkatsu"
	app.Usage = "scrape from mypage.aikatsu.com"
	app.Version = "0.1.0"

	//app.Action =

	app.Commands = []cli.Command{
		{
			Name:     "user",
			Aliases:  []string{"u", "mydata"},
			Category: loginNotRequiredCategoryLabel,
			Usage:    "Fetch Idol data and print as JSON",
			Action:   user,
		},
		{
			Name:     "digital_binder",
			Aliases:  []string{"binder", "db"},
			Category: loginNotRequiredCategoryLabel,
			Usage:    "Fetch Digital Binder data and print as JSON",
			Action:   binder,
		},
		{
			Name:     "card",
			Aliases:  []string{"c"},
			Category: loginNotRequiredCategoryLabel,
			Usage:    "Fetch Card data and print as JSON",
			Action:   card,
		},
		{
			Name: "friends",
			// Aliases:  []string{""},
			Category: loginRequiredCategoryLabel,
			Usage:    "Nothing implemented yet. Sorry!",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
