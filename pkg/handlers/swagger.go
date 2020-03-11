package handlers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
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
				w.Write(d)
			}),
		}}
	return hs, nil
}

func MountSwaggerToGin(path string, router GinRouter) error{
	if ss, err := ServeSwaggerFile(path); err != nil {
		return err
	} else {
		Handlers(ss).MountToEndpoint(router)
		return nil
	}
}
