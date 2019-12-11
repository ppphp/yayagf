package util

import (
	"errors"
	"os"
	"path/filepath"
)

var ErrNoRoot = errors.New("cannot find root")

// find dir contains go.mod
func FindAppRoot(path string) (string, error) {
	if _, err := os.Open(filepath.Join(path, "go.mod")); err == os.ErrNotExist {
		if path != "/" {
			return FindAppRoot(filepath.Dir(path))
		} else {
			return "", errors.New("no project, sb")
		}
	} else if err != nil {
		return "", err
	} else {
		return path, nil
	}
	return "", nil
}

func createDir(dir string) error {
	if err := os.Mkdir(dir, 0744); err == os.ErrNotExist {
		if err := createDir(filepath.Dir(dir)); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func CreateFile(path string, force bool) (*os.File, error) {
	if err := createDir(filepath.Dir(path)); err != nil {
		return nil, err
	}
	o := os.O_RDWR | os.O_CREATE
	if force {
		o |= os.O_TRUNC
	}
	return os.OpenFile(path, o, 0644)
}

func CreateDir(path string, force bool) error {
	return createDir(path)
}
