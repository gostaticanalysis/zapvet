package main

import (
	"github.com/gostaticanalysis/zapvet/passes/fieldtype"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(fieldtype.Analyzer) }

