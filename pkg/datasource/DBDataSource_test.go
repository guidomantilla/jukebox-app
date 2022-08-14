package datasource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDBDataSourceFromDriverName_Ok(t *testing.T) {

	dataSource, err := GetDBDataSourceFromDriverName("mysql", "some_user", "some_password", "some_url")

	assert.NotNil(t, dataSource)
	assert.Nil(t, err)
}

func TestGetDBDataSourceFromDriverName_Error(t *testing.T) {

	dataSource, err := GetDBDataSourceFromDriverName("some_driver_name", "some_user", "some_password", "some_url")

	assert.NotNil(t, err)
	assert.Nil(t, dataSource)
}
