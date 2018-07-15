package fileutils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateUpdateSymlink(t *testing.T) {
	symlink := "testdata/test-symlink"
	defer os.Remove(symlink)
	origFile := "testdata/test-orig.txt"
	origContent, err := ReadFileContentAsString(origFile)
	if err != nil {
		t.Errorf("Could not open file %s", origFile)
	}
	if CreateUpdateSymlink(origFile, symlink) != nil {
		t.Log(err)
		t.Errorf("CreateUpdateSymlink returned error")
	}
	if _, e := os.Stat(symlink); os.IsNotExist(e) {
		t.Errorf("Symlink %s not created", symlink)
	}
	symlinkContent, _ := ReadFileContentAsString(symlink)
	if symlinkContent != origContent {
		t.Errorf(
			"Unexpected content from symlink file\nExpected: %s\nActual: %s",
			origContent,
			symlinkContent)
	}
}

func TestCreateUpdateSymlinkReplacesExistingSymlink(t *testing.T) {
	symlink := "testdata/test-symlink"
	defer os.Remove(symlink)
	origFile := "testdata/test-orig.txt"
	anotherFile := "testdata/test-orig-two.txt"
	expectedContent, err := ReadFileContentAsString(anotherFile)
	if err != nil {
		t.Errorf("Could not open file %s", origFile)
	}
	if err = CreateUpdateSymlink(origFile, symlink); err != nil {
		t.Errorf("CreateUpdateSymlink returned error: %v", err)
	}
	if err = CreateUpdateSymlink(anotherFile, symlink); err != nil {
		t.Errorf("CreateUpdateSymlink returned error: %v", err)
	}
	if _, err = os.Stat(symlink); os.IsNotExist(err) {
		t.Errorf("Symlink %s not created", symlink)
	}
	symlinkContent, _ := ReadFileContentAsString(symlink)
	if symlinkContent != expectedContent {
		t.Errorf(
			"Unexpected content from symlink file\nExpected: %s\nActual: %s",
			expectedContent,
			symlinkContent)
	}
}

func ReadFileContentAsString(file string) (string, error) {
	downloadedData, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(downloadedData[:]), nil
}
