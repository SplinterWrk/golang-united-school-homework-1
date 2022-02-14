package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	res := GetMessage()
	assert.NotEmpty(t, res)
	assert.Equal(t, helloWorld, res)
}
