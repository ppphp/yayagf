package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGinrus(t *testing.T) {
	re := gin.Default()
	g := Ginrus(logrus.New())
	require.NotNil(t, g)

	re.Use(g)
	re.GET("/", func(c *gin.Context) {
		_, _ = ioutil.ReadAll(c.Request.Body)
	})
	re.GET("/bad", func(c *gin.Context) {
		_ = c.Error(fmt.Errorf("aa"))
	})

	t.Run("ok", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, "http://0.0.0.0:8080/", &bytes.Buffer{})
		require.NoError(t, err)
		re.ServeHTTP(rr, r)
		require.NotEqual(t, "", rr.Result())
	})

	t.Run("bad", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, "http://0.0.0.0:8080/bad", &bytes.Buffer{})
		require.NoError(t, err)
		re.ServeHTTP(rr, r)
		require.NotEqual(t, "", rr.Result())
	})

}
