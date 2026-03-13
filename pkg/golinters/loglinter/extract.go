package loglinter

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

func checkAndExtractLogCall(pass *analysis.Pass, call *ast.CallExpr) *LogCall {
	if isSlogCall(pass, call) || isZapCall(pass, call) {
		return extractLogCall(pass, call)
	}
	return nil
}

// isLogCall checks whether the given call expression represents a log method invocation
// from the specified package. It uses type information to verify that the call is a method
// on a named type belonging to the package with the provided path [pkgName]
// and that the method name is among the allowed set
// Returns true if the call matches these criteria, false otherwise
func isLogCall(pass *analysis.Pass,
	call *ast.CallExpr,
	pkgName string,
	allowed map[string]struct{}) bool {

	selExpr, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	obj := pass.TypesInfo.Uses[selExpr.Sel]
	if obj == nil {
		return false
	}

	pkg := obj.Pkg()
	if pkg == nil || pkg.Path() != pkgName {
		return false
	}

	funcObj, ok := obj.(*types.Func)
	if !ok {
		return false
	}
	_, ok = allowed[funcObj.Name()]
	return ok
}

func isSlogCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	allowed := map[string]struct{}{
		"Info":  {},
		"Error": {},
		"Warn":  {},
		"Debug": {},
	}
	return isLogCall(pass, call, "log/slog", allowed)
}

func isZapCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	allowed := map[string]struct{}{
		"Info":  {},
		"Error": {},
		"Warn":  {},
		"Debug": {},
		"Fatal": {},
	}
	return isLogCall(pass, call, "go.uber.org/zap", allowed)
}

func extractLogCall(pass *analysis.Pass, call *ast.CallExpr) *LogCall {
	selExpr := call.Fun.(*ast.SelectorExpr)
	method := selExpr.Sel.Name
	if len(call.Args) > 0 {
		if message, litPos := getStringValue(pass, call.Args[0]); message != "" {
			return &LogCall{
				Pos:      call.Pos(),
				LitPos:   litPos,
				Message:  message,
				Function: method,
			}
		}
	}
	return nil
}

func getStringValue(pass *analysis.Pass, expr ast.Expr) (string, token.Pos) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			val, _ := strconv.Unquote(e.Value)
			return val, e.Pos()
		}
	case *ast.Ident:
		obj := pass.TypesInfo.ObjectOf(e)
		if obj == nil {
			break
		}
		if con, ok := obj.(*types.Const); ok {
			val := con.Val().String()
			if con.Type() == types.Typ[types.String] {
				if len(val) >= 2 && val[0] == '"' && val[len(val)-1] == '"' {
					val = val[1 : len(val)-1]
				}
				return val, e.Pos()
			}
		}

	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			left, pos := getStringValue(pass, e.X)
			right, _ := getStringValue(pass, e.Y)
			return left + right, pos
		}
	}
	return "", 0
}
