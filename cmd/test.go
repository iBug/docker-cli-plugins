package cmd

import (
	"fmt"
	"os"

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
		fmt.Fprintln(DockerCli.Err(), "Error:", err)
		return
	}
	fmt.Fprintln(DockerCli.Out(), exe)
}

func init() {
	RootCmd.AddCommand(TestCmd)
}
