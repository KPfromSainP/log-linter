package loglinter

import (
	"fmt"
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/KPfromSainP/log-linter/pkg/golinters/config"
	"github.com/KPfromSainP/log-linter/pkg/golinters/loglinter/rules"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type LogCall struct {
	Pos      token.Pos
	LitPos   token.Pos
	Message  string
	Function string
}

var Analyzer = &analysis.Analyzer{
	Name:     "loglinter",
	Doc:      "Checks log messages for compliance with rules",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspectResult := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	if len(pass.Files) > 0 {
		currentDir := filepath.Dir(pass.Fset.Position(pass.Files[0].Pos()).Filename)

		for filepath.Dir(currentDir) != currentDir {
			if config.LoadConfig(filepath.Join(currentDir, ".golangci.yml")) == nil {
				break
			}
			currentDir = filepath.Dir(currentDir)
		}
	}

	checkFailed := 0
	inspectResult.Preorder(nodeFilter, func(node ast.Node) {
		call := node.(*ast.CallExpr)
		logCall := checkAndExtractLogCall(pass, call)
		if logCall == nil {
			return
		}
		checkFailed += checkMessage(pass, logCall)
		return
	})
	if checkFailed > 0 {
		pass.Reportf(token.NoPos, "total issues found: %d", checkFailed)
	}
	return nil, nil
}

func makeDiag(call *LogCall, message string) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:     call.LitPos,
		Message: message,
	}
}

func makeSuggestedFix(diag *analysis.Diagnostic, end token.Pos, message string, newText []byte) {
	diag.End = end + 2
	edit := analysis.TextEdit{
		Pos:     diag.Pos + 1,
		End:     end,
		NewText: newText,
	}
	fix := analysis.SuggestedFix{
		Message:   message,
		TextEdits: []analysis.TextEdit{edit},
	}
	diag.SuggestedFixes = []analysis.SuggestedFix{fix}
}

func checkMessage(pass *analysis.Pass, logCall *LogCall) int {
	checksFailed := 0

	if config.CheckRulesConfig("lower-case") && !rules.IsLowercase(logCall.Message) {
		firstRune, runeSize := utf8.DecodeRuneInString(logCall.Message)
		lowerRune := unicode.ToLower(firstRune)
		newText := []byte(string(lowerRune))
		diag := makeDiag(logCall, fmt.Sprintf("log message \"%s\" should start with a lowercase letter (first character: %q)", logCall.Message, firstRune))
		makeSuggestedFix(&diag, logCall.LitPos+1+token.Pos(runeSize), "Change first letter to lowercase", newText)
		pass.Report(diag)
		checksFailed += 1
	}
	if config.CheckRulesConfig("check-english") && !rules.IsEnglish(logCall.Message) {
		diag := makeDiag(logCall, fmt.Sprintf("log message \"%s\" should be in English", logCall.Message))
		pass.Report(diag)
		checksFailed += 1
	}
	if config.CheckRulesConfig("spec-symbols") && !rules.IsNoSpecSymbols(strings.ToLower(logCall.Message)) {
		diag := makeDiag(logCall, fmt.Sprintf("log message \"%s\" should not contain special symbols", logCall.Message))
		newText := []byte(removeSpecialSymbols(logCall.Message))
		makeSuggestedFix(&diag, logCall.LitPos+1+token.Pos(len(logCall.Message)), "Remove spec symbols", newText)
		pass.Report(diag)
		checksFailed += 1
	}
	if config.CheckRulesConfig("sensitive-data") && rules.IsSensitiveData(logCall.Message) {
		diag := makeDiag(logCall, fmt.Sprintf("log message \"%s\" may contains sensitive data (password, token, api_key, etc)", logCall.Message))
		pass.Report(diag)
		checksFailed += 1
	}
	return checksFailed
}

func removeSpecialSymbols(s string) string {
	var result strings.Builder
	for _, r := range s {
		if rules.IsNoSpecSymbol(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}
