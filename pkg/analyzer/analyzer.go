package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type GoImportsAnalyzer struct {
	pattern  string
	Analyzer *analysis.Analyzer
}

func NewGoImportsAnalyzer(pattern string) *analysis.Analyzer {
	a := &GoImportsAnalyzer{
		pattern: pattern,
	}

	return &analysis.Analyzer{
		Name:             "goimportsorder",
		Doc:              "Checks that imports are not mixed with company imports",
		RunDespiteErrors: true,
		Run:              a.run,
		Requires:         []*analysis.Analyzer{inspect.Analyzer},
	}

}

func (a *GoImportsAnalyzer) run(pass *analysis.Pass) (interface{}, error) {

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
			if strings.Contains(imp.Path.Value, a.pattern) && !shouldBeBusinessImport {
				shouldBeBusinessImport = true
				continue
			}
			if !strings.Contains(imp.Path.Value, a.pattern) && shouldBeBusinessImport {
				pass.Reportf(imp.Pos(), "goimportsorder: import %s is not allowed", imp.Path.Value)
				return
			}
		}
	})

	return nil, nil
}
