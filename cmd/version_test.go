package cmd

import (
	"testing"
)

func TestPrintVersion(t *testing.T) {
	version := getVersion()
	if version != "dunix version: 0.0.1-alpha-beta" {
		t.Error("Wrong version returned")
	}
}
