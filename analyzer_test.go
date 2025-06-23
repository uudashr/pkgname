package pkgname_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/uudashr/pkgname"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), pkgname.Analyzer, "a", "b", "c", "d", "e", "f")
}

func TestAnalyzer_IncludeImportAlias(t *testing.T) {
	importAliasAnalyzer := pkgname.NewAnalyzer()

	err := importAliasAnalyzer.Flags.Set("include-import-alias", "true")
	if err != nil {
		t.Fatalf("Failed to set flag: %v", err)
	}

	analysistest.Run(t, analysistest.TestData(), importAliasAnalyzer, "aa")
}
