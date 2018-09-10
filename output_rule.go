package randomstring

import (
	"bytes"
	"strings"
)

// OutputRuleFunc checks if the string meets the rule
type OutputRuleFunc func(str []byte, c byte) bool

// NewBeginWith returns a new output rule func that checks if str does start with a one of letters
func NewBeginWith(letters string) OutputRuleFunc {
	return OutputRuleFunc(func(str []byte, c byte) bool {
		if len(str) > 0 {
			return true
		}
		if strings.IndexByte(letters, c) == -1 {
			return false
		}
		return true
	})
}

// NewNoDuplicateCharacters returns a new output rule func that checks if str doesn't have c
func NewNoDuplicateCharacters() OutputRuleFunc {
	return OutputRuleFunc(func(str []byte, c byte) bool {
		return bytes.IndexByte(str, c) == -1
	})
}

// NewNoSequentialCharacters returns new output rule func that checks if str doesn't have n sequentials characters
func NewNoSequentialCharacters(n uint) OutputRuleFunc {
	return OutputRuleFunc(func(str []byte, c byte) bool {
		lStr := uint(len(str))
		n1 := n - 1
		if lStr < n1 {
			return true
		}
		start := byte(n1)
		valid := false
		for i := lStr - n1; i < lStr; i++ {
			if str[i] != c-start {
				valid = true
				break
			}
			start--
		}
		return valid
	})
}
