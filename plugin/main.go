package plugin

import (
	"github.com/KPfromSainP/log-linter/pkg/golinters/loglinter"
	"golang.org/x/tools/go/analysis"
)

var AnalyzerPlugin = map[string]*analysis.Analyzer{
	"loglinter": loglinter.Analyzer,
}
