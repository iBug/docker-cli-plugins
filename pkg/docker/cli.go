package docker

import "github.com/docker/cli/cli/command"

// Cli is the Docker CLI interface
// It should be populated by the plugin.Run function
var Cli command.Cli
