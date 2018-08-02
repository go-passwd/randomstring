# randomstring

[![GoDoc](https://godoc.org/github.com/tomi77/go-stringgen/randomstring?status.svg)](https://godoc.org/github.com/tomi77/go-stringgen/randomstring)

Go random string generator

## Usage

~~~go
import "github.com/tomi77/go-stringgen/randomstring"
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
