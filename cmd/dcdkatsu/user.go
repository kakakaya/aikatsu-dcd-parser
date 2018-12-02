package main

import (
	"fmt"

	"github.com/kakakaya/aikatsu-dcd-parser"
	"github.com/urfave/cli"
)

func user(c *cli.Context) error {
	fmt.Printf("Hello %q", c.Args().Get(0))

	fmt.Println("Hello friend!")

	var idolID = c.Args().Get(0)
	fmt.Println(idolID)
	idol, err := dcdkatsu.FetchIdol(idolID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("%+v\n", idol)

	return nil
}
