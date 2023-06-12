package update

import (
	"github.com/iBug/docker-ibug/cmd"
	"github.com/iBug/docker-ibug/pkg/updater"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update this tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updater.UpdateBinary()
	},
}

func init() {
	cmd.RootCmd.AddCommand(UpdateCmd)
}
