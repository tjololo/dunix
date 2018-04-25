package unarchive

import (
	"testing"
	"os"
	"io/ioutil"
)

func TestUntarFile(t *testing.T) {
	destination := "testdata/"
	archiveFile := "testdata/test.tar.gz"
	r, err := os.Open(archiveFile)
	if err != nil {
		t.Errorf("File %s not found", archiveFile)
	}
	untaredFile := "testdata/test-untar.txt"
	defer os.Remove(untaredFile)
	Untar(destination, r)
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