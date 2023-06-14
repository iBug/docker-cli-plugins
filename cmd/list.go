package cmd

import (
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/iBug/docker-ibug/pkg/docker"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const TruncateCommandLength = 20

func formatImage(s string) string {
	if strings.HasPrefix(s, "sha256:") {
		return "<none>"
	}
	return s
}

func formatCommand(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.TrimSpace(s)
	if len(s) > TruncateCommandLength {
		s = s[:TruncateCommandLength-1] + "…"
	}
	return s
}

func listContainers(cmd *cobra.Command, args []string) error {
	client := docker.Cli.Client()
	options := types.ContainerListOptions{
		All: true,
	}
	containers, err := client.ContainerList(cmd.Context(), options)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(docker.Cli.Out())
	table.SetCenterSeparator("  ")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetTablePadding("  ")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetNoWhiteSpace(true)

	table.SetHeader([]string{"Name", "Flags", "Image", "Command", "Status"})
	for _, container := range containers {
		flags := docker.GetContainerFlags(container)
		table.Append([]string{
			strings.TrimPrefix(container.Names[0], "/"),
			flags.String(),
			formatImage(container.Image),
			formatCommand(container.Command),
			container.Status,
		})
	}
	table.Render()
	return nil
}

var ListCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List containers",
	Aliases: []string{"list", "ps"},
	RunE:    listContainers,
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
