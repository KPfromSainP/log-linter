package main

import (
	"fmt"
	"os"

	"github.com/KPfromSainP/log-linter/pkg/golinters/loglinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: log-linter <file.go>")
		return
	}
	singlechecker.Main(loglinter.Analyzer)
}
