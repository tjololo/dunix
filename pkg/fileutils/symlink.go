package fileutils

import (
	"os"
	"path/filepath"
)

//CreateUpdateSymlink creates symlink to file or folder. If symlink exists it will be deleted and recreated
func CreateUpdateSymlink(origFile string, symlink string) error {
	absolutePath, err := filepath.Abs(origFile)
	if err != nil {
		return err
	}
	if _, e := os.Stat(symlink); e == nil {
		if os.Remove(symlink) != nil {
			return err
		}
	}
	return os.Symlink(absolutePath, symlink)
}
