//go:build ignore
// +build ignore

package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestWPSFuncPrefix(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "example")
}
