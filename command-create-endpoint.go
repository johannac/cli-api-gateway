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
	"fmt"

	api "github.com/akamai/AkamaiOPEN-edgegrid-golang/api-endpoints-v2"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var flagsCreateEndpoint *api.CreateEndpointOptions = &api.CreateEndpointOptions{}

var commandCreateEndpoint cli.Command = cli.Command{
	Name:        "create-endpoint",
	ArgsUsage:   "",
	Description: "This operation imports an API definition file and creates a new endpoint based on the file contents. You either upload or specify a URL to a Swagger 2.0 or RAML 0.8 file to import details about your API.",
	HideHelp:    true,
	Action:      callCreateEndpoint,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "format",
			Usage:       "Format of the input file, either 'json', 'raml', or 'swagger'",
			Destination: &flagsCreateEndpoint.Format,
		},
		cli.StringFlag{
			Name:        "file",
			Usage:       "Absolute path to the file containing the API definition.",
			Destination: &flagsCreateEndpoint.ImportFile,
		},
		cli.StringFlag{
			Name:        "contract",
			Usage:       "[Swagger or RAML] The unique identifier for the contract under which to provision the endpoint.",
			Destination: &flagsCreateEndpoint.ContractId,
		},
		cli.StringFlag{
			Name:        "group",
			Usage:       "[Swagger or RAML] The unique identifier for the group under which to provision the endpoint.",
			Destination: &flagsCreateEndpoint.GroupId,
		},
	},
}

func callCreateEndpoint(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Creating new API endpoint...",
		fmt.Sprintf("Creating new API endpoint...... [%s]", color.GreenString("OK")),
	)

	endpoint, err := api.CreateEndpoint(flagsCreateEndpoint)
	return output(c, endpoint, err)
}
