package rules

import "unicode"

func IsNoSpecSymbols(message string) bool {
	for _, r := range message {
		if !IsNoSpecSymbol(r) {
			return false
		}
	}
	return true
}

func IsNoSpecSymbol(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r)
}
