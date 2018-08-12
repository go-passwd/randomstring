package randomstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

}

func TestGenerator_Generate(t *testing.T) {
	g, _ := New(NewLength(10), NewIncludeCharset("abc"), NewSimpleGenerate())
	s, err := g.Generate()
	assert.Nil(t, err)
	assert.Len(t, *s, 10)
}

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

func TestNewSimpleGenerate(t *testing.T) {
	rule := NewSimpleGenerate()
	s := rule("abc", 10)
	assert.Len(t, *s, 10)
}
