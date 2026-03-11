package rules

import "unicode"

func IsNoSpecSymbols(message string) bool {
	for _, r := range message {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) || unicode.IsPunct(r)) {
			return false
		}
	}
	return true
}
