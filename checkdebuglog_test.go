package checkdebuglog_test

import (
	"testing"

	"github.com/Yamashou/checkdebuglog"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, checkdebuglog.Analyzer, "a")
}
