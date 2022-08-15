package properties

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PropertiesAdd(t *testing.T) {

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

func Test_PropertiesGet(t *testing.T) {

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

func Test_PropertiesAsMap(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	internalMap := properties.AsMap()

	assert.Equal(t, properties.internalMap, internalMap)
}

func Test_NewDefaultProperties(t *testing.T) {
	properties := NewDefaultProperties()

	assert.NotNil(t, properties)
	assert.NotNil(t, properties.internalMap)
}

func Test_PropertiesBuilderFromArray_Ok(t *testing.T) {

	osArgs := os.Environ()
	properties := NewDefaultProperties(FromArray(&osArgs))

	assert.NotNil(t, properties)
	assert.Equal(t, len(osArgs), len(properties.internalMap))
}

func Test_PropertiesBuilderFromArray_Error(t *testing.T) {

	args := []string{"ola", "adios"}
	properties := NewDefaultProperties(FromArray(&args))

	assert.NotNil(t, properties)
	assert.Equal(t, 0, len(properties.internalMap))
}
