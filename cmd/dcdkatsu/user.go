package main

import (
	"fmt"

	"github.com/kakakaya/aikatsu-dcd-parser"
	"github.com/urfave/cli"
)

func user(c *cli.Context) error {
	var idolID = c.Args().Get(0)
	idol, err := dcdkatsu.FetchIdol(idolID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("%+v\n", idol)
	fmt.Println(idol.AvatarURL)

	return nil
}
