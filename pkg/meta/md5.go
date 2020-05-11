package meta

import (
	"crypto/md5"
	"io/ioutil"
	"os"
)

func CalculateSelfMD5() (string, error) {
	bs, err := ioutil.ReadFile(os.Args[0])
	if err != nil {
		return "", err
	}
	h := md5.New()
	_, err = h.Write(bs)
	if err != nil {
		return "", err
	}
	return string(h.Sum(nil)), nil
}
