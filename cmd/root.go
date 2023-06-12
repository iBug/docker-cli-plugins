package cmd

import (
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ibug",
	Short: "iBug's Docker CLI plugin",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func MakeCmd(dockerCli command.Cli) *cobra.Command {
	return rootCmd
}
