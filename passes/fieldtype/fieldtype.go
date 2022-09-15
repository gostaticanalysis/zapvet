package fieldtype

import (
	"go/ast"
	"go/types"
	"strconv"
	"strings"

	"github.com/gostaticanalysis/zapvet/utils"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "fieldtype finds confliction type of field"

var (
	flagIgnoreFuncs string
)

var Analyzer = &analysis.Analyzer{
	Name: "fieldtype",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func init() {
	Analyzer.Flags.StringVar(&flagIgnoreFuncs, "ignore", "Any", "comma separated ignore function names")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	zappkg := utils.FilterForZap(pass)
	if zappkg == nil {
		return nil, nil
	}
	fs := fieldFuncs(zappkg)
	for _, ignore := range strings.Split(flagIgnoreFuncs, ",") {
		delete(fs, strings.TrimSpace(ignore))
	}

	m := make(map[string]string)
	nodes := []ast.Node{(*ast.CallExpr)(nil)}
	inspect.Preorder(nodes, func(n ast.Node) {
		call, _ := n.(*ast.CallExpr)
		funcname, key := fieldkey(pass, fs, call)
		if fn, exist := m[key]; exist {
			if fn != funcname {
				pass.Reportf(n.Pos(), "%q conflict type %s vs %s", key, funcname, fn)
			}
		} else {
			m[key] = funcname
		}
	})

	return nil, nil
}

func fieldFuncs(zappkg *types.Package) map[string]*types.Func {
	fs := make(map[string]*types.Func)
	scope := zappkg.Scope()
	fieldtyp := scope.Lookup("Field").Type()
	for _, name := range scope.Names() {
		obj, _ := scope.Lookup(name).(*types.Func)
		if obj == nil {
			continue
		}

		sig, _ := obj.Type().(*types.Signature)
		if sig == nil {
			continue
		}

		rets := sig.Results()
		if rets.Len() != 1 ||
			!types.Identical(rets.At(0).Type(), fieldtyp) {
			continue
		}

		params := sig.Params()
		if params.Len() != 2 ||
			!types.Identical(params.At(0).Type(), types.Typ[types.String]) {
			continue
		}

		fs[obj.Name()] = obj
	}
	return fs
}

func fieldkey(pass *analysis.Pass, fs map[string]*types.Func, call *ast.CallExpr) (funcname, key string) {
	var id *ast.Ident
	switch fun := call.Fun.(type) {
	case *ast.Ident:
		id = fun
	case *ast.SelectorExpr:
		id = fun.Sel
	default:
		return "", ""
	}

	obj := pass.TypesInfo.ObjectOf(id)
	if obj == nil {
		return "", ""
	}

	f := fs[obj.Name()]
	if f == nil || f != obj {
		return "", ""
	}

	tv := pass.TypesInfo.Types[call.Args[0]]
	if tv.Value == nil {
		return "", ""
	}

	key, err := strconv.Unquote(tv.Value.String())
	if err != nil {
		return "", ""
	}

	return f.Name(), key
}
