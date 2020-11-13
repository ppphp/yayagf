package handlers

import (
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
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
				_, _ = w.Write(d)
			})},
		)
		return nil
	}); err != nil {
		return nil, err
	}
	return hs, nil
}

func MountStaticHandlerToGin(path string, r gin.IRouter) error {
	if ss, err := ServeStaticDirectory(path); err != nil {
		return err
	} else {
		Handlers(ss).MountToEndpoint(r)
	}
	return nil
}
