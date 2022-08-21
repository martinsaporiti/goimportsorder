package duplicated

import (
	"github.com/martinsaporiti/goimportsorder/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func Something() {
	analyzer := analyzer.NewGoImportsAnalyzer("github.com/martinsaporiti/goimportsorder")
	singlechecker.Main(analyzer)
}
