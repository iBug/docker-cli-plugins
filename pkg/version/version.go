package version

import "github.com/docker/cli/cli-plugins/manager"

var Version string = "devel"

var Metadata = manager.Metadata{
	SchemaVersion: "0.1.0",
	Vendor:        "iBug",
	Version:       Version,
}
