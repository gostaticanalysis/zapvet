package zapvet

import (
	"github.com/gostaticanalysis/zapvet/passes/fieldtype"
	"golang.org/x/tools/go/analysis"
)

// Analyzers returns analyzers of zapvet.
func Analyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		fieldtype.Analyzer,
	}
}
