package randomstring

import "strings"

// CharsetRuleFunc modify a charset and returns it
type CharsetRuleFunc func(charset string) string

// NewIncludeCharset returns a new charset rule func which add chars to charset
func NewIncludeCharset(chars string) CharsetRuleFunc {
	return CharsetRuleFunc(func(charset string) string {
		return charset + chars
	})
}

// NewExcludeCharset returns a new charset rule func which removes chars from charset
func NewExcludeCharset(chars string) CharsetRuleFunc {
	return CharsetRuleFunc(func(charset string) string {
		sCharset := strings.Split(charset, "")
		sChars := strings.Split(chars, "")
		for _, char := range sChars {
			for idx := range sCharset {
				if char == sCharset[idx] {
					sCharset = append(sCharset[:idx], sCharset[idx+1:]...)
					break
				}
			}
		}
		return strings.Join(sCharset, "")
	})
}
