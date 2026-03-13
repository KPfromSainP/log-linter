package rules

import "testing"

func TestIsNoSpecSymbols(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    bool
	}{
		{
			name:    "empty string",
			message: "",
			want:    true,
		},
		{
			name:    "latin letters only",
			message: "HelloWorld",
			want:    true,
		},
		{
			name:    "latin letters with spaces",
			message: "Hello World",
			want:    true,
		},
		{
			name:    "latin letters with digits",
			message: "Hello123",
			want:    true,
		},
		{
			name:    "digits only",
			message: "123456",
			want:    true,
		},
		{
			name:    "spaces only",
			message: "   \t\n",
			want:    true,
		},
		{
			name:    "cyrillic letters",
			message: "Привет",
			want:    true,
		},
		{
			name:    "chinese characters",
			message: "你好",
			want:    true,
		},
		{
			name:    "arabic letters",
			message: "مرحبا",
			want:    true,
		},
		{
			name:    "devanagari digits",
			message: "१२३",
			want:    true,
		},
		{
			name:    "punctuation (ASCII)",
			message: "Hello!",
			want:    false,
		},
		{
			name:    "punctuation (Unicode)",
			message: "Hello…",
			want:    false,
		},
		{
			name:    "emoji",
			message: "Hello 🚀",
			want:    false,
		},
		{
			name:    "currency symbols",
			message: "Price $100",
			want:    false,
		},
		{
			name:    "mathematical symbols",
			message: "a + b = c",
			want:    false,
		},
		{
			name:    "control characters",
			message: "Hello\x00World",
			want:    false,
		},
		{
			name:    "non-breaking space",
			message: "Hello\u00A0World",
			want:    true,
		},
		{
			name:    "mixed valid and invalid",
			message: "Hello, 123",
			want:    false,
		},
		{
			name:    "only invalid characters",
			message: "!@#$%",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNoSpecSymbols(tt.message); got != tt.want {
				t.Errorf("IsNoSpecSymbols(%q) = %v, want %v", tt.message, got, tt.want)
			}
		})
	}
}
