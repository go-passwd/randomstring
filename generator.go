package randomstring

import (
	"bytes"
	"crypto/rand"
	"errors"
	"math/big"
	mrand "math/rand"
	"reflect"
	"strings"
)

// Predefined errors
var (
	errorNoLength  = errors.New("string length must be defined and greater than 0")
	errorNoCharset = errors.New("charset must be set and can't be empty")
	errorNoGen     = errors.New("generator function must be set")
)

// Rules interfaces
type (
	// LengthRuleFunc returns a length of a string to generate
	LengthRuleFunc func() uint

	// CharsetRuleFunc modify a charset and returns it
	CharsetRuleFunc func(charset string) string

	// GenerateRuleFunc generates a new string based on: charset, length and output rules
	GenerateRuleFunc func(charset string, length uint, outputRules []OutputRuleFunc) *string

	// OutputRuleFunc checks if the string meets the rule
	OutputRuleFunc func(str []byte, c byte) bool
)

// NewLength returns a new length rule func which sets string length to n
func NewLength(n uint) LengthRuleFunc {
	return LengthRuleFunc(func() uint {
		return n
	})
}

// NewLengthRange returns a new length rule func which sets string length to length between min and max
func NewLengthRange(min, max uint) LengthRuleFunc {
	return LengthRuleFunc(func() uint {
		return min + uint(mrand.Intn(int(max-min)))
	})
}

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

// NewSimpleGenerate returns a new generate rule func with simple random string generator
func NewSimpleGenerate() GenerateRuleFunc {
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

// Generator is a advanced random string generator based on rules
type Generator struct {
	lengthRule   LengthRuleFunc
	charsetRules []CharsetRuleFunc
	generateRule GenerateRuleFunc
	outputRules  []OutputRuleFunc

	charset string
}

// New creates a new Generator generator
func New(rules ...interface{}) (*Generator, error) {
	g := new(Generator)
	for idx := range rules {
		ruleType := reflect.ValueOf(rules[idx]).Type()
		if ruleType == reflect.TypeOf((*LengthRuleFunc)(nil)).Elem() {
			g.lengthRule = rules[idx].(LengthRuleFunc)
		} else if ruleType == reflect.TypeOf((*CharsetRuleFunc)(nil)).Elem() {
			g.charsetRules = append(g.charsetRules, rules[idx].(CharsetRuleFunc))
		} else if ruleType == reflect.TypeOf((*GenerateRuleFunc)(nil)).Elem() {
			g.generateRule = rules[idx].(GenerateRuleFunc)
		} else if ruleType == reflect.TypeOf((*OutputRuleFunc)(nil)).Elem() {
			g.outputRules = append(g.outputRules, rules[idx].(OutputRuleFunc))
		}
	}
	g.charset = ""
	for idx := range g.charsetRules {
		g.charset = g.charsetRules[idx](g.charset)
	}
	if !reflect.ValueOf(g.lengthRule).IsValid() {
		return nil, errorNoLength
	}
	if len(g.charsetRules) == 0 || g.charset == "" {
		return nil, errorNoCharset
	}
	if !reflect.ValueOf(g.generateRule).IsValid() {
		return nil, errorNoGen
	}
	return g, nil
}

// Generate generates a new random string based of rules
func (g *Generator) Generate() (*string, error) {
	length := g.lengthRule()
	if length == 0 {
		return nil, errorNoLength
	}

	return g.generateRule(g.charset, length, g.outputRules), nil
}
