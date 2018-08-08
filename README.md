# randomstring

[![Build Status](https://travis-ci.org/tomi77/go-stringgen.svg?branch=master)](https://travis-ci.org/tomi77/go-stringgen)
[![Coverage Status](https://coveralls.io/repos/github/tomi77/go-stringgen/badge.svg?branch=master)](https://coveralls.io/github/tomi77/go-stringgen?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomi77/go-stringgen)](https://goreportcard.com/report/github.com/tomi77/go-stringgen)
[![GoDoc](https://godoc.org/github.com/tomi77/go-stringgen?status.svg)](https://godoc.org/github.com/tomi77/go-stringgen)

Go random string generator

## Usage

~~~go
import "goget.in/randomstring.v1"
~~~

### Generate random string

Generating a random 20-character string with lower letters, upper letters and digits:

~~~go
randomstring.Generate(20)
~~~

Generating a random 20-character string with lower letters:

~~~go
randomstring.Generate(20, randomstring.LowerLetters)
~~~

### Predefined character sets

* ``randomstring.LowerLetters`` - lower latin letters
* ``randomstring.UpperLetters`` - upper latin letters
* ``randomstring.Digits`` - digits
