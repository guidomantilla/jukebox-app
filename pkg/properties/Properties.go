package properties

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

var _ Properties = (*DefaultProperties)(nil)

type Properties interface {
	Add(property string, value string)
	Get(property string) string
	AsMap() map[string]string
}

// DefaultProperties

type DefaultProperties struct {
	internalMap map[string]string
}

func (p *DefaultProperties) Add(property string, value string) {
	if p.internalMap[property] == "" {
		p.internalMap[property] = value
	}
}

func (p *DefaultProperties) Get(property string) string {
	return p.internalMap[property]
}

func (p *DefaultProperties) AsMap() map[string]string {
	return p.internalMap
}

//

type DefaultPropertiesOption func(properties *DefaultProperties)

func NewDefaultProperties(options ...DefaultPropertiesOption) *DefaultProperties {
	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}

	for _, opt := range options {
		opt(properties)
	}

	return properties
}

func FromArray(array *[]string) DefaultPropertiesOption {
	return func(properties *DefaultProperties) {
		if array != nil {
			for _, env := range *array {
				pair := strings.SplitN(env, "=", 2)
				if len(pair) != 2 {
					zap.L().Error(fmt.Sprintf("[%s=??] not a key value parameter. expected [key=value]", pair[0]))
					continue
				}
				properties.Add(pair[0], pair[1])
			}
		}
	}
}
