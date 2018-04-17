// +build unix
package idea

import (
	"testing"
	"os"
	"log"
)

func TestGetDownloadURI_version_defined(t *testing.T) {
	expected := "https://download.jetbrains.com/idea/ideaIU-2018.2.tar.gz"
	actual := GetDownloadURI("2018.2")
	if actual != expected {
		t.Error("Wrong download url returned\nExpected:", expected, "\nActual: ", actual)
	}
}

func TestGetDownloadURI_version_not_defined(t *testing.T) {
	expected := "https://download.jetbrains.com/idea/ideaIU-2018.1.tar.gz"
	actual := GetDownloadURI("")
	if actual != expected {
		t.Error("Wrong download url returned\nExpected:", expected, "\nActual: ", actual)
	}
}

func TestDownloadFile(t *testing.T) {
	downloadPath := "./idea.tar.gz"
	//TODO: httptest to get a test not dependent on external dependencies
	DownloadFile(downloadPath, "https://download.jetbrains.com/idea/ideaIU-2018.1.1.tar.gz")
	defer deleteFile(downloadPath)
	var _, e = os.Stat(downloadPath)
	if os.IsNotExist(e) {
		t.Error("File was not downloaded")
	}
}

func deleteFile(file string) {
	log.Printf("Deleting file %s", file)
	os.Remove(file)
}