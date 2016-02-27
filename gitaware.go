package main

import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/codeskyblue/go-sh"

	"fmt"
	"io/ioutil"
)

var pluginCommand = "git-push"

type PushMetadataPlugin struct {
}

func (c *PushMetadataPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] != pluginCommand {
		return
	}
	output, _ := sh.Command("git", "rev-parse", "HEAD").Output()
	fmt.Println("SHA" + string(output))

	err := ioutil.WriteFile(".cfmetadata", output, 0644)
	if err != nil {
		fmt.Println("error " + err.Error())
	}

	args[0] = "push"
	_, err = cliConnection.CliCommand(args...)
	if err != nil {
		fmt.Println("cli error " + err.Error())
	}
	// Pass the args to cf push
}

func (c *PushMetadataPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "GitPushPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     pluginCommand,
				HelpText: "Basic plugin command's help text",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: pluginCommand + "push-metadata\n   cf " + pluginCommand,
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(PushMetadataPlugin))
}
