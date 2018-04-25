package fileutils

import (
	"os"
	"path/filepath"
)

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
