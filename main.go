package main

import (
	"github.com/iBug/docker-ibug/cmd"

	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
)

var metadata = manager.Metadata{
	SchemaVersion: "0.1.0",
	Vendor:        "iBug",
}

func main() {
	plugin.Run(cmd.MakeCmd, metadata)
}
