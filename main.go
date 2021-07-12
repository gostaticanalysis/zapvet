package main

import (
	"github.com/gostaticanalysis/zapvet/zapvet"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(zapvet.Analyzers()...) }
