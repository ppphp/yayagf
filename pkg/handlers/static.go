package handlers

import (
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ServeStaticDirectory(dir string) ([]Handler, error) {
	dir = filepath.Clean(dir)
	hs := []Handler{}
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		d, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		m := mime.TypeByExtension(filepath.Ext(info.Name()))
		hs = append(hs, Handler{
			path: strings.TrimPrefix(path, dir),
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", m)
				w.Write(d)
			})},
		)
		return nil
	}); err != nil {
		return nil, err
	}
	return hs, nil
}
