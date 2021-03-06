package main

import (
	"encoding/json"
	"fmt"

	dcd "github.com/kakakaya/aikatsu-dcd-parser"
	"github.com/urfave/cli"
)

func user(c *cli.Context) error {
	var idolID = c.Args().Get(0)
	idol, err := dcd.FetchIdol(idolID)
	if err != nil {
		return err
	}
	i, err := json.Marshal(idol)
	if err != nil {
		return err
	}
	fmt.Println(string(i))

	return nil
}
