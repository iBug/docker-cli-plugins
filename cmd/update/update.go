package update

import (
	"github.com/iBug/docker-ibug/cmd"
	"github.com/iBug/docker-ibug/pkg/docker"
	"github.com/iBug/docker-ibug/pkg/updater"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update this tool",
	RunE: func(c *cobra.Command, args []string) error {
		return updater.UpdateBinary(docker.Cli.Out())
	},
}

func init() {
	cmd.RootCmd.AddCommand(UpdateCmd)
}
