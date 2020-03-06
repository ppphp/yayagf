package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

func ServeStaticDirectory(path string) ([]Handler, error) {
	path = filepath.Clean(path)
	hs := []Handler{}
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		hs = append(hs, Handler{path: path, handler: http.FileServer(http.Dir(path))})
		return nil
	}); err != nil {
		return nil, err
	}
	return hs, nil
}
