package nopkgvarerrorf

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "nopkgvarerrorf is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "nopkgvarerrorf",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ValueSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if vspec, ok := n.(*ast.ValueSpec); ok {
			for _, value := range vspec.Values {
				if call, ok := value.(*ast.CallExpr); ok {
					if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
						if _, ok := sel.X.(*ast.Ident); ok && sel.Sel.Name == "Errorf" {
							pass.Reportf(n.Pos(), "do not use Errorf at the package level")
						}
					}
					if ident, ok := call.Fun.(*ast.Ident); ok && ident.Name == "Errorf" {
						pass.Reportf(n.Pos(), "do not use Errorf at the package level")
					}
				}
			}
		}
	})

	return nil, nil
}
