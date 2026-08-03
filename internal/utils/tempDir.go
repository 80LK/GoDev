package utils

import (
	"os"
	"path/filepath"
)

func GetTempDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := filepath.Join(cwd, ".tmp")
	if err := os.MkdirAll(dir, 0766); err != nil && !os.IsExist(err) {
		panic(err)
	}

	if err := os.Mkdir(dir, 0766); err != nil && !os.IsExist(err) {
		panic(err)
	}
	return dir
}

func GetTempFile(filename string) string {
	return filepath.Join(GetTempDir(), filename)
}
