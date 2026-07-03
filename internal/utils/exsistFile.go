package utils

import (
	"os"

	"github.com/80LK/godev/errors"
)

func ExsistFile(path string) (bool, error) {
	stat, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if stat.Mode().IsRegular() {
		return true, nil
	} else {
		return false, errors.ErrNotFile
	}
}

func ExsistDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if stat.Mode().IsDir() {
		return true, nil
	} else {
		return false, errors.ErrNotFile
	}
}
