// +build unix
package idea

import (
	"testing"
)

func TestGetDownloadURI(t *testing.T) {
	expected := "https://download.jetbrains.com/idea/ideaIU-2018.1.tar.gz"
	actual := GetDownloadURI()
	if actual != expected {
		t.Error("Wrong download url returned\nExpected:", expected, "\nActual: ", actual)
	}
}