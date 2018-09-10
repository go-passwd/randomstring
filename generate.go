package randomstring

import "strings"

// Generate returns a random string of length n consisting of lower letters, upper letters and digitis
//
// Usage:
//  randomstring.Generate(20, randomstring.Digits, "@#$^&")
//
// Example output:
//  1*&$^6^$#*15&$3427$2
func Generate(n uint, baseChars ...string) (*string, error) {
	letterBytes := strings.Join(baseChars, "")
	if len(letterBytes) == 0 {
		letterBytes = LowerLetters + UpperLetters + Digits
	}
	g, err := New(NewLength(n), NewIncludeCharset(letterBytes))
	if err != nil {
		return nil, err
	}
	return g.Generate()
}
