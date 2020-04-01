package file

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
}

var goModRegexp = regexp.MustCompile("^module (.+)$")

// 不看源码了，internal没法直接import，意思意思得了
func GetMod(path string) (string, error) {
	root, err := FindAppRoot(path)
	if err != nil {
		return "", err
	}
	file, err := ioutil.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(file), "\n") {
		if goModRegexp.MatchString(line) {
			return goModRegexp.FindAllStringSubmatch(line, -1)[0][1], nil
		}
	}
	return "", errors.New("no error")
}
