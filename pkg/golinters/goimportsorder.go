package golinters

import (
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"github.com/quasilyte/go-ruleguard/analyzer"
	"golang.org/x/tools/go/analysis"
)

func NewGoPrintfFuncName() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"goimportsorder",
		"Checks that imports are not mixed with company import",
		[]*analysis.Analyzer{analyzer.Analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
