package fileutils

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"os"
	"io/ioutil"
)

func TestDownloadFile(t *testing.T) {
	downloadPath := "./test-file.txt"
	defer deleteFile(downloadPath)
	expectedFilecontent := "Just a test string"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expectedFilecontent))
	}))
	defer ts.Close()
	DownloadFile(downloadPath, ts.URL)
	var _, e = os.Stat(downloadPath)
	if os.IsNotExist(e) {
		t.Error("File was not downloaded")
	}
	downloadedData, _ := ioutil.ReadFile(downloadPath)
	contentString := string(downloadedData[:])
	if contentString != expectedFilecontent {
		t.Errorf(
			"Unexpected content of downloaded file\nExpected: %s\nActual: %s",
			expectedFilecontent,
			contentString)
	}
}

func deleteFile(file string) {
	os.Remove(file)
}
