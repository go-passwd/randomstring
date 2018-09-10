package randomstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNoDuplicateCharacters(t *testing.T) {
	rule := NewNoDuplicateCharacters()
	assert.False(t, rule([]byte("abc"), byte('a')))
	assert.False(t, rule([]byte("abc"), byte('b')))
	assert.False(t, rule([]byte("abc"), byte('c')))
	assert.True(t, rule([]byte("abc"), byte('d')))
}

func TestNewBeginWith(t *testing.T) {
	rule := NewBeginWith("abc")
	assert.True(t, rule([]byte("abc"), byte('c')))
	assert.True(t, rule([]byte("a"), byte('a')))
	assert.True(t, rule([]byte("a"), byte('d')))
	assert.True(t, rule([]byte(""), byte('c')))
	assert.False(t, rule([]byte(""), byte('d')))
}

func TestNewNoSequentialCharacters(t *testing.T) {
	rule := NewNoSequentialCharacters(3)
	assert.True(t, rule([]byte("a"), byte('b')))
	assert.True(t, rule([]byte("ab"), byte('d')))
	assert.False(t, rule([]byte("ab"), byte('c')))
	assert.False(t, rule([]byte("xab"), byte('c')))
	assert.True(t, rule([]byte("xab"), byte('d')))
}
