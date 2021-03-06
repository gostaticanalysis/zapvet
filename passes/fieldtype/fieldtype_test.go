package fieldtype_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/gostaticanalysis/zapvet/passes/fieldtype"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	fieldtype.Analyzer.Flags.Set("ignore", "Any, Reflect")
	analysistest.Run(t, testdata, fieldtype.Analyzer, "a")
}
