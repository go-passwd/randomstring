package randomstring

import "math/rand"

// LengthRuleFunc returns a length of a string to generate
type LengthRuleFunc func() uint

// NewLength returns a new length rule func which sets string length to n
func NewLength(n uint) LengthRuleFunc {
	return LengthRuleFunc(func() uint {
		return n
	})
}

// NewLengthRange returns a new length rule func which sets string length to length between min and max
func NewLengthRange(min, max uint) LengthRuleFunc {
	return LengthRuleFunc(func() uint {
		return min + uint(rand.Intn(int(max-min)))
	})
}
