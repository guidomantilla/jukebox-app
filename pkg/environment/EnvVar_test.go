package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnvVar(t *testing.T) {

	aux := "some string"
	evnVar := NewEnvVar(aux)
	assert.Equal(t, aux, evnVar.AsString())
}

func TestEnvVarAsString(t *testing.T) {

	aux := "some string"
	evnVar := NewEnvVar(aux)
	assert.Equal(t, aux, evnVar.AsString())
}

func TestEnvVarAsInt_Ok(t *testing.T) {

	aux := "101010"
	evnVar := NewEnvVar(aux)
	value, _ := evnVar.AsInt()
	assert.Equal(t, 101010, value)
}

func TestEnvVarAsInt_Error(t *testing.T) {

	aux := "some string"
	evnVar := NewEnvVar(aux)
	value, _ := evnVar.AsInt()
	assert.Equal(t, 0, value)
}
