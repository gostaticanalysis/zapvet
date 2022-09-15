package utils

import (
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)

func FilterForZap(pass *analysis.Pass) *types.Package {
	for _, pkg := range pass.Pkg.Imports() {
		if analysisutil.RemoveVendor(pkg.Path()) == "go.uber.org/zap" {
			return pkg
		}
	}
	return nil
}
