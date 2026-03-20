package main

import (
	"github.com/KPfromSainP/log-linter/pkg/golinters/loglinter"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{loglinter.Analyzer}, nil
}
