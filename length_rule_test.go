package randomstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLength(t *testing.T) {
	rule := NewLength(10)
	length := rule()
	assert.Equal(t, length, uint(10))
}

func TestNewLengthRange(t *testing.T) {
	rule := NewLengthRange(10, 20)
	length := rule()
	assert.NotNil(t, length)
	comp := func() bool {
		return length >= 10 && length <= 20
	}
	assert.Condition(t, comp)
}
