package main

import (
	"github.com/cloudfoundry/cli/plugin"

	"fmt"
)

type PushMetadataPlugin struct {
}

func (c *PushMetadataPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	for arg := range args {
		fmt.Println(arg)
	}
	// Do your logic.
	// Pass the args to cf push
}

func (c *PushMetadataPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "PushMetadataPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "push-metadata",
				HelpText: "Basic plugin command's help text",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "push-metadata\n   cf push-metadata",
				},
			},
		},
	}
}

func main () {
	plugin.Start(new(PushMetadataPlugin))
}
