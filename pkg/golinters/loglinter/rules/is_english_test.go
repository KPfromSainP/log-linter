package rules

import "testing"

func TestIsEnglish(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    bool
	}{
		{
			name:    "pure ascii letters",
			message: "HelloWorld",
			want:    true,
		},
		{
			name:    "ascii letters with spaces",
			message: "Hello World",
			want:    true,
		},
		{
			name:    "ascii letters with punctuation",
			message: "Hello, world!",
			want:    true,
		},
		{
			name:    "ascii letters with digits",
			message: "Hello123",
			want:    true, // цифры входят в ASCII 32-127
		},
		{
			name:    "only punctuation",
			message: "!?.,;:-",
			want:    true,
		},
		{
			name:    "unicode punctuation (en dash)",
			message: "Hello – world", // en dash (U+2013) – считается пунктуацией по Unicode
			want:    true,            // unicode.IsPunct(U+2013) == true
		},
		{
			name:    "unicode punctuation (ellipsis)",
			message: "Hello…world", // горизонтальное многоточие (U+2026) – пунктуация
			want:    true,
		},
		{
			name:    "cyrillic letters",
			message: "Привет",
			want:    false,
		},
		{
			name:    "chinese characters",
			message: "你好",
			want:    false,
		},
		{
			name:    "emoji",
			message: "Hello 🚀",
			want:    false,
		},
		{
			name:    "newline",
			message: "Hello\nworld",
			want:    true,
		},
		{
			name:    "newline",
			message: "Hello\tworld",
			want:    true,
		},
		{
			name:    "mixed ascii and cyrillic",
			message: "Hello Привет",
			want:    false,
		},
		{
			name:    "empty string",
			message: "",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEnglish(tt.message); got != tt.want {
				t.Errorf("IsEnglish(%q) = %v, want %v", tt.message, got, tt.want)
			}
		})
	}
}
