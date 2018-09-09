package randomstring

// Character sets allowed in a generated string
const (
	// Set of lowercase latin letters
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// Set of uppercase latin letters
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Set of digits
	Digits = "1234567890"

	// Set of symbols
	Symbols = "!\";#$%&'()*+,-./:;<=>?@[]^_`{|}~"
)

// Character sets disallowed in generated string
const (
	// Set of letters, digits and symbols which are looks similar
	Similar = "il1Lo0O"

	// Set of symbols which are ambigous
	Ambigous = "{}[]()/\\'\"`~,;:.<>"
)
