package zapvet_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gostaticanalysis/zapvet/zapvet"
	"golang.org/x/tools/go/packages"
)

func TestAnalyzers(t *testing.T) {
	want := analyzerNames(t)
	got := zapvet.Analyzers()
	gotNames := make([]string, len(got))
	for i := range got {
		gotNames[i] = got[i].Name
	}
	if diff := cmp.Diff(want, gotNames); diff != "" {
		t.Error(diff)
	}
}

func analyzerNames(t *testing.T) []string {
	t.Helper()
	cfg := &packages.Config{Mode: packages.NeedName}
	pkgs, err := packages.Load(cfg, "github.com/gostaticanalysis/zapvet/passes/...")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	names := make([]string, len(pkgs))
	for i := range pkgs {
		names[i] = pkgs[i].Name
	}

	return names
}
