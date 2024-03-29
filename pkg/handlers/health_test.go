package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestMountHealthHandlerToGin(t *testing.T) {
	g := gin.Default()
	MountHealthHandlerToGin(g.Group("health"))

	req, _ := http.NewRequest("GET", "/health/", nil)
	resp := httptest.NewRecorder()
	g.ServeHTTP(resp, req)

	require.NotNil(t, resp.Body.String())
}

func TestHealthHandlerToGin(t *testing.T) {

	req, _ := http.NewRequest("GET", "/health/", nil)
	resp := httptest.NewRecorder()
	healthHandler[0].handler.ServeHTTP(resp, req)
	require.NotNil(t, healthHandler[0].GetHTTPHandler())

	require.NotNil(t, resp.Body.String())
}
