package randomstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_Generate(t *testing.T) {
	g, _ := New(NewLength(10), NewIncludeCharset("abc"), NewDefaultGenerate())
	s, err := g.Generate()
	assert.Nil(t, err)
	assert.Len(t, *s, 10)
}
