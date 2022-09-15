package ensurekv

import (
	"go/ast"

	"github.com/gostaticanalysis/zapvet/utils"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "ensurekv ensures logger methods are called with key value pairs"

var Analyzer = &analysis.Analyzer{
	Name: "ensurekv",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var sugaredLoggerFunctions = []string{"Errorw", "Infow", "Warnw", "Debugw"}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if utils.FilterForZap(pass) == nil {
		return nil, nil
	}
	nodes := []ast.Node{(*ast.CallExpr)(nil)}
	inspector.Preorder(nodes, func(n ast.Node) {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}
		fun, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		if !functionHasKVs(fun.Sel.Name) {
			return
		}
		if len(call.Args)%2 == 0 {
			pass.Reportf(n.Pos(), "%s needs to be called with a message, key and value, missing some args", fun.Sel.Name)
		}
	})
	return nil, nil

}

func functionHasKVs(name string) bool {
	for _, funcs := range sugaredLoggerFunctions {
		if name == funcs {
			return true
		}
	}
	return false
}
