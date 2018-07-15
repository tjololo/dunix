package cmd

import (
	"testing"
)

func TestPrintVersion(t *testing.T) {
	version := getVersion()
	if version != "dunix version: 0.2.0" {
		t.Error("Wrong version returned")
	}
}
