package file

import (
	"errors"
	"os"
	"path/filepath"
)

var ErrNoRoot = errors.New("cannot find root")

func GetAppRoot() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return FindAppRoot(pwd)
}

// find dir contains go.mod
func FindAppRoot(path string) (string, error) {
	if _, err := os.Open(filepath.Join(path, "go.mod")); os.IsNotExist(err) {
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
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := createDir(filepath.Dir(dir)); err != nil {
			return err
		}
		if err := os.Mkdir(dir, 0755); err != nil {
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

func CreateFileWithContent(path string, content string) error {
	f, err := CreateFile(path, false)
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func CreateDir(path string, force bool) error {
	return createDir(path)
}
