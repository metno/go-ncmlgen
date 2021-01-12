/*
  Copyright 2020 MET Norway

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/metno/go-ncmlgen/pkg/ncml"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ncmlgen",
		Usage: "Create a ncml file aggregating provided NetCDF files",

		Commands: []*cli.Command{
			{
				Name:    "aggregate",
				Aliases: []string{"a"},
				Usage:   "Create a ncml file aggregating provided files",
				Action: func(ctx *cli.Context) error {
					fileList := ctx.Args().Slice()
					ncmlString, err := ncml.CreateNcMLWithAggregation(fileList)
					if err != nil {
						return fmt.Errorf("%v", err)
					}
					fmt.Println(*ncmlString)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
