package randomstring

// Character sets allowed in a password
const (
	// LowerLetters is a sets of lowercase latin letters
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is a set of uppercase latin letters
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is a set of digits
	Digits = "1234567890"

	// Symbols is a set of symbols
	Symbols = "!\";#$%&'()*+,-./:;<=>?@[]^_`{|}~"
)

// Character sets disallowed
const (
	// Similar is a set of letters, digits and symbols which are looks similar
	Similar = "il1Lo0O"

	// Ambigous is a set of sumbols which are ambigous
	Ambigous = "{}[]()/\\'\"`~,;:.<>"
)
