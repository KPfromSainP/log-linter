package rules

import (
	"unicode"
	"unicode/utf8"
)

func IsLowercase(message string) bool {
	if len(message) == 0 {
		return true
	}
	r, _ := utf8.DecodeRuneInString(message)
	return !unicode.IsUpper(r)
}
