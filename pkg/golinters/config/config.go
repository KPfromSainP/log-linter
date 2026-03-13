package config

import (
	"fmt"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Rules    RulesConfig    `yaml:"rules"`
	Patterns PatternsConfig `yaml:"patterns"`
}

type RulesConfig struct {
	CheckLowerCase     bool `yaml:"lower-case"`
	CheckEnglish       bool `yaml:"check-english"`
	CheckSpecSymbols   bool `yaml:"spec-symbols"`
	CheckSensitiveData bool `yaml:"sensitive-data"`
}

type PatternsConfig struct {
	Patterns          StringSet `yaml:"patterns"`
	SensitiveKeywords StringSet `yaml:"sensitive-keywords"`
	CompiledPatterns  []*regexp.Regexp
}

var cfg Config

func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			// use user log via context or custom Logger interface
		}
	}(file)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	for pattern := range cfg.Patterns.Patterns {
		for keyword := range cfg.Patterns.SensitiveKeywords {
			cfg.Patterns.CompiledPatterns = append(cfg.Patterns.CompiledPatterns,
				buildPattern(keyword, pattern))
		}
	}

	return nil
}

func buildPattern(keyword string, pattern string) *regexp.Regexp {
	escaped := regexp.QuoteMeta(keyword)
	regexPattern := fmt.Sprintf(pattern, escaped)
	return regexp.MustCompile(regexPattern)
}

func GetPatterns() []*regexp.Regexp {
	return cfg.Patterns.CompiledPatterns
}

func CheckRulesConfig(ruleName string) bool {
	switch ruleName {
	case "lower-case":
		return cfg.Rules.CheckLowerCase
	case "spec-symbols":
		return cfg.Rules.CheckSpecSymbols
	case "check-english":
		return cfg.Rules.CheckEnglish
	case "sensitive-data":
		return cfg.Rules.CheckSensitiveData
	default:
		return false
	}
}
