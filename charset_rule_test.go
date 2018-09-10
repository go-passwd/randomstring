package randomstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIncludeCharset(t *testing.T) {
	rule := NewIncludeCharset("bc")
	chars := rule("a")
	assert.Equal(t, "abc", chars)
}

func TestNewExcludeCharset(t *testing.T) {
	rule := NewExcludeCharset("bc")
	chars := rule("abcd")
	assert.Equal(t, "ad", chars)
}
