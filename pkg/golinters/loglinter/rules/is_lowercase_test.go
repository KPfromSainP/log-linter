package rules

import "testing"

func TestIsLowercase(t *testing.T) {
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
			name:    "single lowercase latin letter",
			message: "a",
			want:    true,
		},
		{
			name:    "single uppercase latin letter",
			message: "A",
			want:    false,
		},
		{
			name:    "lowercase latin word",
			message: "hello",
			want:    true,
		},
		{
			name:    "uppercase latin word",
			message: "Hello",
			want:    false,
		},
		{
			name:    "all caps word",
			message: "WORLD",
			want:    false,
		},
		{
			name:    "lowercase cyrillic",
			message: "привет",
			want:    true,
		},
		{
			name:    "uppercase cyrillic",
			message: "Привет",
			want:    false,
		},
		{
			name:    "chinese characters (no case)",
			message: "你好",
			want:    true,
		},
		{
			name:    "arabic (no case)",
			message: "مرحبا",
			want:    true,
		},
		{
			name:    "digits first",
			message: "123abc",
			want:    true,
		},
		{
			name:    "punctuation first",
			message: "!hello",
			want:    true,
		},
		{
			name:    "space first",
			message: " hello",
			want:    true,
		},
		{
			name:    "emoji first",
			message: "🚀 hello",
			want:    true,
		},
		{
			name:    "uppercase with diacritic (Latin)",
			message: "École",
			want:    false,
		},
		{
			name:    "lowercase with diacritic",
			message: "école",
			want:    true,
		},
		{
			name:    "invalid UTF-8 first byte",
			message: string([]byte{0xFF, 'a', 'b', 'c'}),
			want:    true,
		},
		{
			name:    "null byte first",
			message: "\x00abc",
			want:    true,
		},
		{
			name:    "greek uppercase",
			message: "Γειά",
			want:    false,
		},
		{
			name:    "greek lowercase",
			message: "γειά",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLowercase(tt.message); got != tt.want {
				t.Errorf("IsLowercase(%q) = %v, want %v", tt.message, got, tt.want)
			}
		})
	}
}
