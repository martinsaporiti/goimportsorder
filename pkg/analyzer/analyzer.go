package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var (
	Analyzer = &analysis.Analyzer{
		Name:             "goimportsorder",
		Doc:              "Checks that imports are not mixed with company imports",
		RunDespiteErrors: true,
		Run:              run,
		Requires:         []*analysis.Analyzer{inspect.Analyzer},
	}

	pattern string
)

func init() {
	if pattern == "" {
		Analyzer.Flags.StringVar(&pattern, "pattern", "", "pattern to match")
	}
}

func run(pass *analysis.Pass) (interface{}, error) {

	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		fileSpec, ok := node.(*ast.File)
		if !ok {
			return
		}
		shouldBeBusinessImport := false
		for _, imp := range fileSpec.Imports {
			if strings.Contains(imp.Path.Value, pattern) && !shouldBeBusinessImport {
				shouldBeBusinessImport = true
				continue
			}
			if !strings.Contains(imp.Path.Value, pattern) && shouldBeBusinessImport {
				pass.Reportf(imp.Pos(), "goimportsorder: import %s is not allowed", imp.Path.Value)
				return
			}
		}
	})

	return nil, nil
}
