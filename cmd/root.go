package cmd

import (
	"github.com/docker/cli/cli/command"
	"github.com/iBug/docker-ibug/pkg/docker"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ibug",
	Short: "iBug's Docker CLI plugin",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// MakeCmd implements the first argument for docker/cli/cli-plugins/plugin.Run
func MakeCmd(cli command.Cli) *cobra.Command {
	docker.Cli = cli
	RootCmd.SetIn(cli.In())
	RootCmd.SetOut(cli.Out())
	RootCmd.SetErr(cli.Err())
	return RootCmd
}
