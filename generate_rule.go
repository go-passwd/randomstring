package randomstring

import (
	"crypto/rand"
	"math/big"
)

// GenerateRuleFunc generates a new string based on: charset, length and output rules
type GenerateRuleFunc func(charset string, length uint, outputRules []OutputRuleFunc) *string

// NewDefaultGenerate returns a new generate rule func with default random string generator
func NewDefaultGenerate() GenerateRuleFunc {
	return GenerateRuleFunc(func(charset string, length uint, outputRules []OutputRuleFunc) *string {
		b := make([]byte, length)
		letterBytesLength := big.NewInt(int64(len(charset)))
		for i := range b {
			for {
				idx, _ := rand.Int(rand.Reader, letterBytesLength)
				c := charset[idx.Int64()]
				valid := true
				for _, outputRule := range outputRules {
					if !outputRule(b, c) {
						valid = false
						break
					}
				}
				if valid {
					b[i] = c
					break
				}
			}
		}

		s := string(b)
		return &s
	})
}
