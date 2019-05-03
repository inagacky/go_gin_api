package util


import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {

	envName := "GO_API_UTIL_TEST"
	dVal := "DefaultValue"
	sVal := "SetValue"
	v := Getenv(envName, dVal)
	assert.Equal(t, v, dVal, "not Matched DefaultValue")


	os.Setenv(envName, sVal)
	v = Getenv(envName, dVal)
	assert.Equal(t, v, sVal, "not Matched Env Value")

}