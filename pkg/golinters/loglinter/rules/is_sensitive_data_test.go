package rules

import (
	"testing"

	"github.com/KPfromSainP/log-linter/pkg/golinters/config"
)

func TestIsSensitiveData(t *testing.T) {
	err := config.LoadConfig("../../../../.golangci.yml")
	if err != nil {
		return
	}
	patterns := config.GetPatterns()
	if len(patterns) == 0 {
		t.Skip("No patterns loaded, skipping test")
	}

	tests := []struct {
		name    string
		message string
		want    bool
	}{
		{
			name:    "password with colon",
			message: "password: mysecret",
			want:    true,
		},
		{
			name:    "password with equals",
			message: "password=mysecret",
			want:    true,
		},
		{
			name:    "password with spaces",
			message: "password : mysecret",
			want:    true,
		},
		{
			name:    "uppercase PASSWORD",
			message: "PASSWORD: mysecret",
			want:    true,
		},
		{
			name:    "mixed case PassWord",
			message: "PassWord=mysecret",
			want:    true,
		},
		{
			name:    "token with colon",
			message: "token: abc123",
			want:    true,
		},
		{
			name:    "api_key with equals",
			message: "api_key=12345",
			want:    true,
		},
		{
			name:    "secret in middle",
			message: "my_secret: value",
			want:    true,
		},
		{
			name:    "password without colon or equals",
			message: "password is secret",
			want:    false,
		},
		{
			name:    "password with colon no space",
			message: "password:secret",
			want:    true,
		},
		{
			name:    "password as part of word",
			message: "mypassword: secret",
			want:    true, // подстрока "password" найдена
		},
		{
			name:    "multiple lines",
			message: "line1\npassword: secret\nline2",
			want:    true,
		},
		{
			name:    "no sensitive data",
			message: "Hello world",
			want:    false,
		},
		{
			name:    "empty string",
			message: "",
			want:    false,
		},
		{
			name:    "only special characters",
			message: "!@#$%",
			want:    false,
		},
		{
			name:    "password with special chars after colon",
			message: "password: !@#$%",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSensitiveData(tt.message); got != tt.want {
				t.Errorf("IsSensitiveData(%q) = %v, want %v", tt.message, got, tt.want)
			}
		})
	}
}
