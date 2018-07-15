package golang

import "testing"

func TestGetDownloadURI_version_defined(t *testing.T) {
	expected := "https://dl.google.com/go/go1.9.0.linux-amd64.tar.gz"
	actual := GetDownloadURI("1.9.0")
	if actual != expected {
		t.Error("Wrong download url returned\nExpected:", expected, "\nActual: ", actual)
	}
}

func TestGetDownloadURI_version_not_defined(t *testing.T) {
	expected := "https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz"
	actual := GetDownloadURI("")
	if actual != expected {
		t.Error("Wrong download url returned\nExpected:", expected, "\nActual: ", actual)
	}
}

