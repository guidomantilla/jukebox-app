package environment

import (
	"jukebox-app/pkg/properties"
)

var _ Environment = (*DefaultEnvironment)(nil)

type Environment interface {
	GetValue(property string) EnvVar
	GetValueOrDefault(property string, defaultValue string) EnvVar
	GetPropertySources() []properties.PropertySource
	AppendPropertySources(propertySources ...properties.PropertySource)
}

// DefaultEnvironment

type DefaultEnvironment struct {
	propertySources []properties.PropertySource
}

func (environment *DefaultEnvironment) GetValue(property string) EnvVar {

	var value string
	for _, source := range environment.propertySources {
		internalValue := source.Get(property)
		if internalValue != "" {
			value = internalValue
			break
		}
	}
	return NewEnvVar(value)
}

func (environment *DefaultEnvironment) GetValueOrDefault(property string, defaultValue string) EnvVar {

	envVar := environment.GetValue(property)
	if envVar != "" {
		return envVar
	}
	return NewEnvVar(defaultValue)
}

func (environment *DefaultEnvironment) GetPropertySources() []properties.PropertySource {
	return environment.propertySources
}

func (environment *DefaultEnvironment) AppendPropertySources(propertySources ...properties.PropertySource) {
	environment.propertySources = append(environment.propertySources, propertySources...)
}

//

type DefaultEnvironmentOption = func(environment *DefaultEnvironment)

func NewDefaultEnvironment(options ...DefaultEnvironmentOption) *DefaultEnvironment {
	environment := &DefaultEnvironment{
		propertySources: make([]properties.PropertySource, 0),
	}
	for _, opt := range options {
		opt(environment)
	}

	return environment
}

func WithPropertySources(propertySources ...properties.PropertySource) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		environment.propertySources = propertySources
	}
}
