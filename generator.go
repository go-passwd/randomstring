package randomstring

import (
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
	// LengthRuleFunc defined a type of length rule func
	LengthRuleFunc func() uint

	// CharsetRuleFunc defined a type of charset rule func
	CharsetRuleFunc func(charset string) string

	// GenerateRuleFunc defined a type of generate rule func
	GenerateRuleFunc func(charset string, length uint) *string
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

// NewSimpleGenerate returns a new generate rule func with simple random string generator
func NewSimpleGenerate() GenerateRuleFunc {
	return GenerateRuleFunc(func(charset string, length uint) *string {
		b := make([]byte, length)
		letterBytesLength := big.NewInt(int64(len(charset)))
		for i := range b {
			idx, _ := rand.Int(rand.Reader, letterBytesLength)
			b[i] = charset[idx.Int64()]
		}

		s := string(b)
		return &s
	})
}

// Generator is a advanced random string generator based on rules
type Generator struct {
	LengthRule   LengthRuleFunc
	CharsetRules []CharsetRuleFunc
	GenerateRule GenerateRuleFunc
}

// New creates a new Generator generator
func New(rules ...interface{}) (*Generator, error) {
	g := new(Generator)
	for idx := range rules {
		ruleType := reflect.ValueOf(rules[idx]).Type()
		if ruleType == reflect.TypeOf((*LengthRuleFunc)(nil)).Elem() {
			g.LengthRule = rules[idx].(LengthRuleFunc)
		} else if ruleType == reflect.TypeOf((*CharsetRuleFunc)(nil)).Elem() {
			g.CharsetRules = append(g.CharsetRules, rules[idx].(CharsetRuleFunc))
		} else if ruleType == reflect.TypeOf((*GenerateRuleFunc)(nil)).Elem() {
			g.GenerateRule = rules[idx].(GenerateRuleFunc)
		}
	}
	if !reflect.ValueOf(g.LengthRule).IsValid() {
		return nil, errorNoLength
	}
	if len(g.CharsetRules) == 0 {
		return nil, errorNoCharset
	}
	if !reflect.ValueOf(g.GenerateRule).IsValid() {
		return nil, errorNoGen
	}
	return g, nil
}

// Generate generates a new random string based of rules
func (g *Generator) Generate() (*string, error) {
	length := g.LengthRule()
	chars := ""
	for idx := range g.CharsetRules {
		chars = g.CharsetRules[idx](chars)
	}
	if length == 0 {
		return nil, errorNoLength
	}
	if chars == "" {
		return nil, errorNoCharset
	}

	return g.GenerateRule(chars, length), nil
}
