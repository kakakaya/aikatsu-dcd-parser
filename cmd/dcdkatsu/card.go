package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kakakaya/aikatsu-dcd-parser"
	"github.com/urfave/cli"
)

func card(c *cli.Context) error {
	var cardURL = c.Args().Get(0)
	_, err := url.ParseRequestURI(cardURL)
	if err != nil {
		return err
	}

	card, err := dcd.FetchCard(cardURL)
	if err != nil {
		return err
	}
	i, err := json.Marshal(card)
	if err != nil {
		return err
	}
	fmt.Println(string(i))

	return nil
}
