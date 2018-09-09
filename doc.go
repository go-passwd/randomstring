// Package randomstring provides:
//
// - function `Generate` to generate random string based on set of characters
//
// - type `Generator` to generate random string based on set of rules
//
// Example
//
// This code:
//
//  gen := randomstring.New(randomstring.NewLength(20), randomstring.NewInlcudeCharset(randomstring.LowerLetters), randomstring.NewNoDuplicateCharacters(), randomstring.NewSimpleGenerator())
//  fmt.Println(gen.Generate())
//
// will generate random string that:
//
// - has length 20
//
// - contains lowercase latin letters
//
// - all letters in random string are unique
//
// Custom rules
//
// Generator accept four types of rules:
//
// - Length rule that defines a random string length
//
// - Charset rule that build a character set
//
// - Output rule that checks built string
//
// - Generator rule that generates a random string
//
// All types can be customized, ex:
//
//  func NewCustomCharset() CharsetRuleFunc {
//    return CharsetRuleFunc(func(charset string) string {
//      return "abc"
//    })
//  }
//
// Usage:
//
//  gen := randomstring.New(randomstring.NewLength(20), NewCustomCharset(), randomstring.NewNoDuplicateCharacters(), randomstring.NewSimpleGenerator())
package randomstring
