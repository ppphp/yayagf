package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gitlab.papegames.com/fengche/yayagf/pkg/prom"
)

func TestMountPromHandlerToGin(t *testing.T) {
	g := gin.Default()
	MountPromHandlerToGin(g.Group("metrics"), prom.SysCPU())

	req, _ := http.NewRequest("GET", "/metrics/", nil)
	resp := httptest.NewRecorder()
	g.ServeHTTP(resp, req)

	require.NotNil(t, resp.Body.String())
}
