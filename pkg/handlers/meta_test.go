package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestMountMetaHandlerToGin(t *testing.T) {
	g := gin.Default()
	MountMetaHandlerToGin(g.Group("meta"))

	req, _ := http.NewRequest("GET", "/meta/", nil)
	resp := httptest.NewRecorder()
	g.ServeHTTP(resp, req)

	require.NotNil(t, resp.Body.String())
}
