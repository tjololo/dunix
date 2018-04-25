package fileutils

import (
	"io"
	"compress/gzip"
	"log"
	"archive/tar"
	"path/filepath"
	"os"
)
const UNTAR_PERMISSION = 0755

func Untar(destination string, archive io.Reader) error {
	gzr, err := gzip.NewReader(archive)
	defer gzr.Close()
	if err != nil {
		log.Printf("Could not read gzip")
		return err
	}
	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		switch {

		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}

		target := filepath.Join(destination, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.Mkdir(target, UNTAR_PERMISSION); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			file, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return nil
			}
			defer file.Close()

			if _, err := io.Copy(file, tr); err != nil {
				return err
			}
		}
	}
}
