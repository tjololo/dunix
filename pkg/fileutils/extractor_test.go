package fileutils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestUntarFile(t *testing.T) {
	destination := "testdata/"
	archiveFile := "testdata/test.tar.gz"
	untaredFile := "testdata/test-untar.txt"
	defer os.Remove(untaredFile)
	Untar(archiveFile, destination)
	_, e := os.Stat(untaredFile)
	if os.IsNotExist(e) {
		t.Errorf("Expected %s to be present after untar %s", untaredFile, archiveFile)
	}
	untaredData, _ := ioutil.ReadFile(untaredFile)
	contentString := string(untaredData[:])
	expectedFilecontent := "Hello I was in a tar.gz file\n"
	if contentString != expectedFilecontent {
		t.Errorf(
			"Unexpected content of downloaded file\nExpected: {%s}\nActual: {%s}",
			expectedFilecontent,
			contentString)
	}
}
