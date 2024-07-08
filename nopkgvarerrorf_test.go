package nopkgvarerrorf_test

import (
	"testing"

	"nopkgvarerrorf"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, nopkgvarerrorf.Analyzer, "a")
}
