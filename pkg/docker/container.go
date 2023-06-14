package docker

import (
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/iBug/docker-ibug/pkg/namesgenerator"
)

type ContainerFlags struct {
	IsAutoName bool
}

func GetContainerFlags(c types.Container) ContainerFlags {
	return ContainerFlags{
		IsAutoName: namesgenerator.IsAutoName(strings.TrimPrefix(c.Names[0], "/")),
	}
}

func (c ContainerFlags) String() string {
	fIsAutoName := "-"
	if c.IsAutoName {
		fIsAutoName = "A"
	}
	return fIsAutoName
}
