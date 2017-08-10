package normalizers

import (
	"strings"

	"github.com/sebdah/trollhunter/languages"
)

type Simple struct{}

// Constructor
func NewSimple() *Simple {
	return &Simple{}
}

// Basic normalizer. Main functions:
// - Trim leading and trailing whitespaces
// - Convert to lower case
// - Remove punctuation and commas
func (this *Simple) Normalize(data string, language *languages.Language) string {
	// Trim leading and trailing whitespaces
	data = strings.TrimSpace(data)

	// Convert to lower case
	data = strings.ToLower(data)

	// Remove punctuation and commas
	log.Debug(string(language.GetAlphabet()))
	log.Info(data)

	return data
}
