package environment

import (
	properties2 "jukebox-app/pkg/properties"
	"os"
)

const (
	CMD_PROPERTY_SOURCE_NAME = "CMD"
	OS_PROPERTY_SOURCE_NAME  = "OS"
)

func LoadEnvironment(cmdArgs *[]string) Environment {

	cmdSource := properties2.NewDefaultPropertySource(CMD_PROPERTY_SOURCE_NAME, properties2.NewDefaultProperties().FromArray(cmdArgs).Build())
	env := NewDefaultEnvironment().WithPropertySources(cmdSource).Build()

	osArgs := os.Environ()
	osSource := properties2.NewDefaultPropertySource(OS_PROPERTY_SOURCE_NAME, properties2.NewDefaultProperties().FromArray(&osArgs).Build())
	env.AppendPropertySources(osSource)

	return env
}
