package languages

type Language interface {
	GetAlphabet() []byte
}

// Return the alphabet for this language
func GetAlphabet(lang Language) []byte {
	return lang.Alphabet
}
