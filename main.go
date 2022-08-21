package main

import (
	"flag"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/martinsaporiti/goimportsorder/pkg/analyzer"
)

func main() {
	pattern := flag.String("pattern", "github.com/foo", "a string")
	flag.Parse()

	analyzer := analyzer.NewGoImportsAnalyzer(*pattern)
	singlechecker.Main(analyzer)
}
