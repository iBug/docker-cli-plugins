package cmd

import (
	"fmt"
	"os"

	"github.com/iBug/docker-ibug/pkg/docker"
	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:    "test",
	Short:  "test command",
	Run:    RunTest,
	Hidden: true,
}

func RunTest(cmd *cobra.Command, args []string) {
	exe, err := os.Executable()
	if err != nil {
		fmt.Fprintln(docker.Cli.Err(), "Error:", err)
		return
	}
	fmt.Fprintln(docker.Cli.Out(), exe)
}

func init() {
	RootCmd.AddCommand(TestCmd)
}
