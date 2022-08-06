package properties

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPropertiesAdd(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	assert.Equal(t, "value01", properties.internalMap["property01"])
	assert.Equal(t, "value02", properties.internalMap["property02"])
	assert.Equal(t, "value03", properties.internalMap["property03"])

}

func TestPropertiesGet(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	assert.Equal(t, "value01", properties.Get("property01"))
	assert.Equal(t, "value02", properties.Get("property02"))
	assert.Equal(t, "value03", properties.Get("property03"))
}

func TestPropertiesAsMap(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	internalMap := properties.AsMap()

	assert.Equal(t, properties.internalMap, internalMap)
}

func TestNewDefaultProperties(t *testing.T) {
	propertiesBuilder := NewDefaultProperties()

	assert.NotNil(t, propertiesBuilder)
	assert.NotNil(t, propertiesBuilder.defaultProperties)
}

func TestPropertiesBuilderFromArray_Ok(t *testing.T) {

	osArgs := os.Environ()
	propertiesBuilder := NewDefaultProperties()
	propertiesBuilder.FromArray(&osArgs)

	assert.NotNil(t, propertiesBuilder)
	assert.NotNil(t, propertiesBuilder.defaultProperties)
	assert.Equal(t, len(osArgs), len(propertiesBuilder.defaultProperties.internalMap))
}

func TestPropertiesBuilderFromArray_Error(t *testing.T) {

	args := []string{"ola", "adios"}
	propertiesBuilder := NewDefaultProperties()
	propertiesBuilder.FromArray(&args)

	assert.NotNil(t, propertiesBuilder)
	assert.NotNil(t, propertiesBuilder.defaultProperties)
	assert.Equal(t, 0, len(propertiesBuilder.defaultProperties.internalMap))
}

func TestPropertiesBuilderBuild(t *testing.T) {

	osArgs := os.Environ()
	propertiesBuilder := NewDefaultProperties()
	properties := propertiesBuilder.FromArray(&osArgs).Build()

	assert.NotNil(t, properties)
}
