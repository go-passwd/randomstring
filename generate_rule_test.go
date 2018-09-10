package randomstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultGenerate(t *testing.T) {
	rule := NewDefaultGenerate()
	s := rule("abc", 10, []OutputRuleFunc{})
	assert.Len(t, *s, 10)
}
