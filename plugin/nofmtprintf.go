// This must be package main
package main

import (
	linter "github.com/fes300/go-lint-no-fmt-print"
	"golang.org/x/tools/go/analysis"
)

type AnalyzerPlugin interface {
	GetAnalyzers() []*analysis.Analyzer
}

type analyzerPlugin struct{}

var _ AnalyzerPlugin = analyzerPlugin{}

func (a analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return linter.Analyzers
}

func New(conf any) ([]*analysis.Analyzer, error) {
	return linter.Analyzers, nil
}
