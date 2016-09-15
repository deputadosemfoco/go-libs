package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureBuildReturnsNotNil(t *testing.T) {
	v := Build()

	assert.NotNil(t, v)
}

func TestNotEmptyWithData(t *testing.T) {
	data := "data"
	validator := Build()
	res := validator.NotEmpty(data, "message")

	val, err := validator.Validate()

	assert.True(t, res)
	assert.Nil(t, err)
	assert.True(t, val)
}

func TestNotEmptyWithoutData(t *testing.T) {
	data := ""
	validator := Build()
	res := validator.NotEmpty(data, "message")

	val, err := validator.Validate()

	assert.False(t, res)
	assert.False(t, val)
	assert.Equal(t, "message", err[0])
}

func TestIsGtInt(t *testing.T) {
	validator := Build()
	res := validator.IsGt(10, 5, "message")
	val, err := validator.Validate()

	assert.True(t, res)
	assert.Nil(t, err)
	assert.True(t, val)
}
