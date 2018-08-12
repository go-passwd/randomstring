package randomstring

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	s, err := Generate(100, Digits, LowerLetters)
	assert.Nil(t, err)
	assert.Len(t, *s, 100)
	for _, l := range strings.Split(UpperLetters, "") {
		assert.NotContains(t, *s, l)
	}

	s, err = Generate(20, Digits)
	assert.Nil(t, err)
	assert.Len(t, *s, 20)
	for _, l := range strings.Split(UpperLetters+LowerLetters, "") {
		assert.NotContains(t, *s, l)
	}
}

func TestGenerate_Default(t *testing.T) {
	s, err := Generate(30)
	assert.Nil(t, err)
	assert.Len(t, *s, 30)
}
