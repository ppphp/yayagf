package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var healthHandler = Handlers([]Handler{{
	path: "",
	handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("health"))
	})}})

func MountHealthHandlerToGin(r gin.IRouter)  {
	healthHandler.MountToEndpoint(r)
}
