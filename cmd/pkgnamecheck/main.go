package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/uudashr/pkgname"
)

func main() {
	singlechecker.Main(pkgname.Analyzer)
}
