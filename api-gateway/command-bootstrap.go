// Copyright 2018. Akamai Technologies, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	api "github.com/akamai/AkamaiOPEN-edgegrid-golang/api-endpoints-v2"
	api2 "github.com/akamai/AkamaiOPEN-edgegrid-golang/apikey-manager-v1"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

var commandLocator akamai.CommandLocator = func() ([]cli.Command, error) {
	commands := []cli.Command{
		cli.Command{
			Name:        "list",
			Description: "List commands",
			Action:      akamai.CmdList,
		},
		cli.Command{
			Name:         "help",
			Description:  "Displays help information",
			ArgsUsage:    "[command] [sub-command]",
			Action:       akamai.CmdHelp,
			BashComplete: akamai.DefaultAutoComplete,
		},
		commandCreateEndpoint,
		commandImportEndpoint,
		commandUpdateEndpoint,
		commandListEndpoints,
		commandListResources,
		commandActivateEndpoint,
		commandRemoveEndpoint,
		commandPrivacy,
		commandClone,
		commandStatus,
	}

	return commands, nil
}

func initConfig(c *cli.Context) error {
	config, err := akamai.GetEdgegridConfig(c)
	if err != nil {
		return err
	}
	api.Init(config)
	api2.Init(config)
	return nil
}

func output(c *cli.Context, toReturn interface{}, err error) error {
	if err != nil {
		akamai.StopSpinnerFail()
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	if c.Bool("json") {
		if toReturn != nil {
			returnJSON, err := json.MarshalIndent(toReturn, "", "  ")
			if err != nil {
				akamai.StopSpinnerFail()
				return cli.NewExitError(color.RedString(err.Error()), 1)
			}

			akamai.StopSpinnerOk()
			fmt.Fprintln(c.App.Writer, string(returnJSON))
			return nil
		}

		akamai.StopSpinnerOk()
		return nil
	}

	akamai.StopSpinnerOk()

	t, ok := toReturn.(Tabular)
	if ok {
		tt := t.ToTable()
		tt.Render()
	}

	return nil

}

func hasSTDIN() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}

	return false
}

type Tabular interface {
	ToTable() *tablewriter.Table
}
