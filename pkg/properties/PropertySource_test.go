package properties

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultPropertySource(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	propertySource := NewDefaultPropertySource("group", properties)

	assert.NotNil(t, propertySource)
	assert.Equal(t, properties, propertySource.properties)
	assert.Equal(t, "group", propertySource.name)
}

func TestPropertySourceGet(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	propertySource := NewDefaultPropertySource("group", properties)

	value := propertySource.Get("property01")

	assert.Equal(t, "value01", value)
}

func TestPropertySourceAsMap(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	propertySource := NewDefaultPropertySource("group", properties)

	internalMap := propertySource.AsMap()

	assert.Equal(t, "group", internalMap["name"])
	assert.Equal(t, properties.internalMap, internalMap["value"])
}
