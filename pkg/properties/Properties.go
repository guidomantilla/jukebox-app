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

// DefaultPropertiesBuilder

type DefaultPropertiesBuilder struct {
	defaultProperties *DefaultProperties
}

func NewDefaultProperties() *DefaultPropertiesBuilder {
	return &DefaultPropertiesBuilder{
		defaultProperties: &DefaultProperties{
			internalMap: make(map[string]string),
		},
	}
}

func (builder *DefaultPropertiesBuilder) Build() *DefaultProperties {
	return builder.defaultProperties
}

func (builder *DefaultPropertiesBuilder) FromArray(array *[]string) *DefaultPropertiesBuilder {

	if array != nil {
		for _, env := range *array {
			pair := strings.SplitN(env, "=", 2)
			if len(pair) != 2 {
				zap.L().Error(fmt.Sprintf("[%s=??] not a key value parameter. expected [key=value]", pair[0]))
				continue
			}
			builder.defaultProperties.Add(pair[0], pair[1])
		}
	}

	return builder
}
