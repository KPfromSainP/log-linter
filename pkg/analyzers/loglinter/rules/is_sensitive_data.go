package rules

import (
	"github.com/KPfromSainP/log-linter/pkg/analyzers/config"
)

func ContainsSensitiveData(message string) bool {
	for _, re := range config.GetPatterns() {
		if re.MatchString(message) {
			return true
		}
	}
	return false
}
