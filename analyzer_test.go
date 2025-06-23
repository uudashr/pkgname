package pkgname_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/uudashr/pkgname"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), pkgname.Analyzer, "a", "b", "c", "d", "e", "f")
}
