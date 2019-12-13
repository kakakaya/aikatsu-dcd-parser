package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kakakaya/aikatsu-dcd-parser"
	"github.com/urfave/cli"
)

func binder(c *cli.Context) error {
	if c.NArg() != 2 {
		return fmt.Errorf("Digital_binder requires exactly two arguments: <OwnerID, SeriesID>, but got %d", c.NArg())
	}
	var ownerID = c.Args().Get(0)
	var seriesID, err = strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return err
	}

	db, err := dcd.FetchDigitalBinder(ownerID, seriesID)
	if err != nil {
		return err
	}
	i, err := json.Marshal(db)
	if err != nil {
		return err
	}
	fmt.Println(string(i))

	return nil
}
