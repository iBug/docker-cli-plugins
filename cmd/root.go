package cmd

import (
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ibug",
	Short: "iBug's Docker CLI plugin",
	Long:  "iBug's Docker CLI plugin, but there's no feature implemented yet.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// DockerCli is the Docker CLI interface
var DockerCli command.Cli

// MakeCmd implements the first argument for docker/cli/cli-plugins/plugin.Run
func MakeCmd(dockerCli command.Cli) *cobra.Command {
	DockerCli = dockerCli
	return RootCmd
}
