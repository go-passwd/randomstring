package randomstring

import (
	"errors"
	"reflect"
)

// Predefined errors
var (
	errorNoLength  = errors.New("string length must be defined and greater than 0")
	errorNoCharset = errors.New("charset must be set and can't be empty")
)

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
	if g.generateRule == nil {
		g.generateRule = NewDefaultGenerate()
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
