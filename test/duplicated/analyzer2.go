package duplicated

import (
	"github.com/martinsaporiti/goimportsorder/pkg/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func Something1() {
	analyzer := analyzer.Analyzer
	singlechecker.Main(analyzer)

}
