package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewEnvVar(t *testing.T) {

	aux := "some string"
	evnVar := NewEnvVar(aux)
	assert.Equal(t, aux, evnVar.AsString())
}

func Test_EnvVarAsString(t *testing.T) {

	aux := "some string"
	evnVar := NewEnvVar(aux)
	assert.Equal(t, aux, evnVar.AsString())
}

func Test_EnvVarAsInt_Ok(t *testing.T) {

	aux := "101010"
	evnVar := NewEnvVar(aux)
	value, _ := evnVar.AsInt()
	assert.Equal(t, 101010, value)
}

func Test_EnvVarAsInt_Error(t *testing.T) {

	aux := "some string"
	evnVar := NewEnvVar(aux)
	value, _ := evnVar.AsInt()
	assert.Equal(t, 0, value)
}
