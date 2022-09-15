package zapvet

import (
	"github.com/gostaticanalysis/zapvet/passes/ensurekv"
	"github.com/gostaticanalysis/zapvet/passes/fieldtype"
	"golang.org/x/tools/go/analysis"
)

// Analyzers returns analyzers of zapvet.
func Analyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		ensurekv.Analyzer,
		fieldtype.Analyzer,
	}
}
