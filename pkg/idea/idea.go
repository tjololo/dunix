package idea

import (
	"os"
	"net/http"
	"io"
	"gopkg.in/cheggaaa/pb.v1"
)

const VERSION = "2018.1.2"

const URI = "https://download.jetbrains.com/idea/ideaIU-"

func GetDownloadURI(version string) string {
	if version == "" {
		return URI + VERSION + FILE_EXTENSION
	}
	return URI + version + FILE_EXTENSION
}

func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bar := pb.New(int(resp.ContentLength)).SetUnits(pb.U_BYTES)
	bar.Start()
	r := bar.NewProxyReader(resp.Body)
	// Write the body to file
	_, err = io.Copy(out, r)
	if err != nil {
		return err
	}

	return nil
}
