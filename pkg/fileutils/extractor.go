package fileutils

import (
	"compress/gzip"
	"archive/tar"
	"os"
	"fmt"
	"github.com/mholt/archiver"
	"strings"
)

//Untar an tar.gz archive to destination folder
func Untar(src string, destination string) error {
	return archiver.TarGz.Open(src, destination)
}

//GetRootFolder root folder in tar.gz folder
func GetRootFolder(src string) string {
	file, err := os.Open(src)
	if err != nil {
		fmt.Printf("Could not open")
	}
	gzr, err := gzip.NewReader(file)
	defer gzr.Close()
	if err != nil {
		fmt.Printf("Could not read gzip")
	}
	tr := tar.NewReader(gzr)
	header, err:= tr.Next()
	if err != nil {
		fmt.Printf("Could not get header")
	}
	firstRecord := header.Name
	index := strings.Index(firstRecord, "/")
	return firstRecord[0:index]
}
