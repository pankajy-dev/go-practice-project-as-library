package stringutilpankaj

import (
	"strings"
	"unicode"
)

// Capitalize converts the first character of a string to uppercase
func Capitalize(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// CapitalizeWords converts the first character of each word to uppercase
func CapitalizeWords(s string) string {
	return strings.Title(strings.ToLower(s))
}

// ToUpper converts all characters in a string to uppercase
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower converts all characters in a string to lowercase
func ToLower(s string) string {
	return strings.ToLower(s)
}
