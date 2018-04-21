// +build unix
package idea

import (
	"testing"
	"os"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
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
