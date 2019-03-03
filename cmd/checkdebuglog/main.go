package main

import (
	"github.com/Yamashou/checkdebuglog"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func main() {
	multichecker.Main(
		inspect.Analyzer,
		checkdebuglog.Analyzer,
	)
}
