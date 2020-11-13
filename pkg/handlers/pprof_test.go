package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestMountPProfHandlerToGin(t *testing.T) {
	g := gin.Default()
	MountPProfHandlerToGin(g.Group("pprof"))

	req, _ := http.NewRequest("GET", "/pprof/", nil)
	resp := httptest.NewRecorder()
	g.ServeHTTP(resp, req)

	require.NotNil(t, resp.Body.String())
}
