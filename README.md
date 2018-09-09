# Go random string generator

[![Build Status](https://travis-ci.org/go-passwd/randomstring.svg?branch=master)](https://travis-ci.org/go-passwd/randomstring)
[![Coverage Status](https://coveralls.io/repos/github/go-passwd/randomstring/badge.svg?branch=master)](https://coveralls.io/github/go-passwd/randomstring?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-passwd/randomstring)](https://goreportcard.com/report/github.com/go-passwd/randomstring)
[![GoDoc](https://godoc.org/github.com/go-passwd/randomstring?status.svg)](https://godoc.org/github.com/go-passwd/randomstring)

## Usage

~~~go
import "github.com/go-passwd/randomstring"
~~~

## Documentation

https://go-passwd.github.io/randomstring.html

## Rules functions

### Length rule function

Returns a length of a string to generate.

#### NewLength

Sets string length to n.

#### NewLengthRange

Sets string length to length between min and max

### Charset rule function

Modify a charset and returns it.

#### NewIncludeCharset

Add chars to charset.

#### NewExcludeCharset

Removes chars from charset.

### Output rule function

#### NewBeginWith

Checks if newly selected at random char does start with a one of letters.
Function only executed at first char.

#### NewNoDuplicateCharacters

Checks if string doesn't have newly selected at random char.

#### NewNoSequentialCharacters

Checks if string doesn't have n sequentials characters.

### Generate rule function

Generates a new string based on: charset, length and output rules.

#### NewSimpleGenerate

Simple random string generator.
