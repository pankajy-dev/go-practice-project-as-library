package stringutilpankaj

import "testing"

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single char lowercase", "a", "A"},
		{"single char uppercase", "A", "A"},
		{"lowercase word", "hello", "Hello"},
		{"uppercase word", "HELLO", "HELLO"},
		{"mixed case", "hELLO", "HELLO"},
		{"with spaces", "hello world", "Hello world"},
		{"unicode", "über", "Über"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Capitalize(tt.input)
			if result != tt.expected {
				t.Errorf("Capitalize(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCapitalizeWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single word", "hello", "Hello"},
		{"two words", "hello world", "Hello World"},
		{"multiple words", "the quick brown fox", "The Quick Brown Fox"},
		{"all caps", "HELLO WORLD", "Hello World"},
		{"mixed case", "hELLo WoRLd", "Hello World"},
		{"with punctuation", "hello, world!", "Hello, World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CapitalizeWords(tt.input)
			if result != tt.expected {
				t.Errorf("CapitalizeWords(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"lowercase word", "hello", "HELLO"},
		{"uppercase word", "HELLO", "HELLO"},
		{"mixed case", "HeLLo", "HELLO"},
		{"with spaces", "hello world", "HELLO WORLD"},
		{"with numbers", "hello123", "HELLO123"},
		{"with punctuation", "hello, world!", "HELLO, WORLD!"},
		{"unicode", "über", "ÜBER"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToUpper(tt.input)
			if result != tt.expected {
				t.Errorf("ToUpper(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"lowercase word", "hello", "hello"},
		{"uppercase word", "HELLO", "hello"},
		{"mixed case", "HeLLo", "hello"},
		{"with spaces", "HELLO WORLD", "hello world"},
		{"with numbers", "HELLO123", "hello123"},
		{"with punctuation", "HELLO, WORLD!", "hello, world!"},
		{"unicode", "ÜBER", "über"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToLower(tt.input)
			if result != tt.expected {
				t.Errorf("ToLower(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}