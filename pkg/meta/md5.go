package meta

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
)

func CalculateSelfMD5() (string, error) {
	bs, err := ioutil.ReadFile(os.Args[0])
	if err != nil {
		return "", err
	}
	h := md5.New()
	_, _ = h.Write(bs) // md5 write never return err
	return hex.EncodeToString(h.Sum(nil)), nil
}
