package randomstring

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// Generate returns a random string of length n consisting of lower letters, upper letters and digitis
//
// Usage:
//  randomstring.Generate(20, randomstring.Digits, "@#$^&")
//
// Example output:
//  1*&$^6^$#*15&$3427$2
func Generate(n int, baseChars ...string) string {
	var letterBytes string
	letterBytes = strings.Join(baseChars, "")
	if len(letterBytes) == 0 {
		letterBytes = LowerLetters + UpperLetters + Digits
	}
	b := make([]byte, n)
	letterBytesLength := big.NewInt(int64(len(letterBytes)))
	for i := range b {
		idx, _ := rand.Int(rand.Reader, letterBytesLength)
		b[i] = letterBytes[idx.Int64()]
	}

	return string(b)
}
