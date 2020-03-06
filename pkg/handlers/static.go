package handlers

import (
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
		hs = append(hs, Handler{path: strings.TrimPrefix(path, dir), handler: http.FileServer(http.Dir(path))})
		return nil
	}); err != nil {
		return nil, err
	}
	return hs, nil
}
