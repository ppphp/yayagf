package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var healthHandler = Handlers([]Handler{{
	path: "",
	handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("health"))
	})}})

func MountHealthHandlerToGin(r gin.IRouter) {
	healthHandler.MountToEndpoint(r)
}
