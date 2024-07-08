package main

import (
	"nopkgvarerrorf"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(nopkgvarerrorf.Analyzer) }
