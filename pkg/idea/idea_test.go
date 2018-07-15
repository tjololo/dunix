package idea

import (
	"testing"
)

func TestGetDownloadURI_version_defined(t *testing.T) {
	expected := "https://download.jetbrains.com/idea/ideaIU-2018.2.tar.gz"
	actual := GetDownloadURI("2018.2")
	if actual != expected {
		t.Error("Wrong download url returned\nExpected:", expected, "\nActual: ", actual)
	}
}

func TestGetDownloadURI_version_not_defined(t *testing.T) {
	expected := "https://download.jetbrains.com/idea/ideaIU-2018.1.2.tar.gz"
	actual := GetDownloadURI("")
	if actual != expected {
		t.Error("Wrong download url returned\nExpected:", expected, "\nActual: ", actual)
	}
}
