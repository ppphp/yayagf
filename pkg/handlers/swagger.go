package handlers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// serve on prefix/swagger.json
func ServeSwaggerFile(swagger string) ([]Handler, error) {
	swagger = filepath.Clean(swagger)
	d, err := ioutil.ReadFile(swagger)
	if err != nil {
		return nil, err
	}
	hs := []Handler{{
		path: "swagger.json",
		handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			_, _ = w.Write(d)
		}),
	}}
	return hs, nil
}

func MountSwaggerFileToGin(path string, router gin.IRouter) error {
	if ss, err := ServeSwaggerFile(path); err != nil {
		return err
	} else {
		Handlers(ss).MountToEndpoint(router)
		return nil
	}
}

func MountSwaggerStringToGin(swagger string, router gin.IRouter) {
	Handlers([]Handler{{
		path: "swagger.json",
		handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			_, _ = w.Write([]byte(swagger))
		}),
	}}).MountToEndpoint(router)
}
