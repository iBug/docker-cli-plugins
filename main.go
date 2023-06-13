package main

import (
	"github.com/iBug/docker-ibug/cmd"
	_ "github.com/iBug/docker-ibug/cmd/update"
	"github.com/iBug/docker-ibug/pkg/version"

	"github.com/docker/cli/cli-plugins/plugin"
)

func main() {
	plugin.Run(cmd.MakeCmd, version.Metadata)
}
