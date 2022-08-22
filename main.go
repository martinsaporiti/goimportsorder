package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/martinsaporiti/goimportsorder/pkg/analyzer"
)

func main() {
	analyzer := analyzer.Analyzer
	singlechecker.Main(analyzer)
}
