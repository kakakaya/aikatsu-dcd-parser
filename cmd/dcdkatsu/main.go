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
			Usage:    "add a task to the list",
			Action:   user,
		},
		{
			Name:     "digital_binder",
			Aliases:  []string{"dbe"},
			Category: loginNotRequiredCategoryLabel,
			Usage:    "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name: "friends",
			// Aliases:  []string{""},
			Category: loginRequiredCategoryLabel,
			Usage:    "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name: "friends",
			// Aliases:  []string{""},
			Category: loginRequiredCategoryLabel,
			Usage:    "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name: "friends",
			// Aliases:  []string{""},
			Category: loginRequiredCategoryLabel,
			Usage:    "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name: "friends",
			// Aliases:  []string{""},
			Category: loginRequiredCategoryLabel,
			Usage:    "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name: "friends",
			// Aliases:  []string{""},
			Category: loginRequiredCategoryLabel,
			Usage:    "complete a task on the list",
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
