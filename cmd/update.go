package cmd

import (
	"github.com/iBug/docker-ibug/pkg/docker"
	"github.com/iBug/docker-ibug/pkg/updater"
	"github.com/spf13/cobra"
)

var fForce bool

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update this tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updater.UpdateBinary(docker.Cli.Out(), fForce)
	},
}

func init() {
	RootCmd.AddCommand(UpdateCmd)

	flags := UpdateCmd.Flags()
	flags.BoolVarP(&fForce, "force", "f", false, "Force update")
}
