package checkdebuglog

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "checkdebuglog",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "checkdebuglog is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if call, ok := n.(*ast.CallExpr); ok {
			var id *ast.SelectorExpr
			var idx *ast.Ident

			switch fun := call.Fun.(type) {
			case *ast.SelectorExpr:
				id = fun
			default:
				return
			}

			switch x := id.X.(type) {
			case *ast.Ident:
				idx = x
			default:
				return

			}

			if id != nil && idx != nil && !pass.TypesInfo.Types[id].IsType() {
				if id.Sel.Name == "Debugf" && idx.Name == "log" {
					pass.Reportf(call.Lparen, "detect log.Debugf use debug code")
				}
			}

		}

	})

	return nil, nil
}
