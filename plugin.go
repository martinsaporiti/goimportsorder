package main

import (
	"github.com/martinsaporiti/goimportsorder/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin
