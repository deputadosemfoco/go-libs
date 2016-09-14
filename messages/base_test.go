package messages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	res := ResBuilder().AsSuccess(nil).Build()

	assert.Equal(t, true, res.Success)
}

func TestError(t *testing.T) {
	res := ResBuilder().AsErr("test").Build()

	assert.Equal(t, false, res.Success)
	assert.Equal(t, "test", res.Message)
}

func TestValidationMessages(t *testing.T) {
	msgs := []string{"1", "2"}
	res := ResBuilder().AsErr("test").WithValMsgs(msgs)

	assert.Equal(t, false, res.Success)
	assert.Equal(t, "test", res.Message)
	assert.Equal(t, 2, len(res.ValidationMessages))
	assert.Equal(t, "1", res.ValidationMessages[0])
	assert.Equal(t, "2", res.ValidationMessages[1])
}
