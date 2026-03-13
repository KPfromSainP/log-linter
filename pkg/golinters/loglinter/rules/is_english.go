package rules

import "unicode"

func IsEnglish(message string) bool {
	for _, r := range message {
		if !isBasicLatin(r) {
			return false
		}
	}
	return true
}

func isBasicLatin(r rune) bool {
	return r >= 32 && r <= unicode.MaxASCII || unicode.IsPunct(r) || unicode.IsSpace(r)
}
