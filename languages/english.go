package languages

type English struct {
	Alphabet []byte
}

func NewEnglish() *English {
	return &English{
		Alphabet: []byte("abcdefghijklmnopqrstuvwzyx"),
	}
}
