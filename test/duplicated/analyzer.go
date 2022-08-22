package duplicated

import (
	"github.com/martinsaporiti/goimportsorder/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func Something() {
	analyzer := analyzer.Analyzer
	singlechecker.Main(analyzer)
}
