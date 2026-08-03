package utils

import "path/filepath"

func GetTempDir() string {
	return "./.tmp"
}

func GetTempFile(filename string) string {
	return filepath.Join(GetTempDir(), filename)
}
