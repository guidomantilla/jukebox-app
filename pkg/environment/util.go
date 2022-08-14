package environment

import (
	"jukebox-app/pkg/properties"
	"os"
)

const (
	CMD_PROPERTY_SOURCE_NAME = "CMD"
	OS_PROPERTY_SOURCE_NAME  = "OS"
)

func LoadEnvironment(cmdArgs *[]string) Environment {

	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OS_PROPERTY_SOURCE_NAME,
		properties.NewDefaultProperties(properties.FromArray(&osArgs)))

	cmdSource := properties.NewDefaultPropertySource(CMD_PROPERTY_SOURCE_NAME,
		properties.NewDefaultProperties(properties.FromArray(cmdArgs)))

	return NewDefaultEnvironment(WithPropertySources(osSource, cmdSource))
}
