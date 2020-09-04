package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.papegames.com/fengche/yayagf/pkg/meta"
)

// 简单的meta信息暴露

func MountMetaHandlerToGin(r gin.IRouter) {
	Handlers{Handler{
		path: "/",
		handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			m, _ := json.Marshal(meta.Get())
			w.Write(m)
		})}}.MountToEndpoint(r)
}
