package cmd

import (
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ibug",
	Short: "iBug's Docker CLI plugin",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// MakeCmd implements the first argument for docker/cli/cli-plugins/plugin.Run
func MakeCmd(dockerCli command.Cli) *cobra.Command {
	return RootCmd
}
