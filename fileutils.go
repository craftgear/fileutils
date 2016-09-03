package fileutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Cp(dest, source string, mode os.FileMode) error {
	b, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dest, b, mode)
}

func CpDir(destDir, sourceDir string, mode os.FileMode) error {
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		relPath, _ := filepath.Rel(sourceDir, path)
		if relPath == "." {
			return nil
		}
		if info.IsDir() {
			return os.MkdirAll(filepath.Join(destDir, relPath), mode)
		} else {
			return cp(filepath.Join(destDir, relPath), path, mode)
		}
	})
}
