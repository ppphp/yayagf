package maotai

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	assert "github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}

func TestDefault(t *testing.T) {
	assert.NotNil(t, Default("test"))
}

func TestNikkiSerializer(t *testing.T) {
	d := Default("test")
	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)
	c.Request, _ = http.NewRequest(http.MethodGet, "http://0.0.0.0", nil)
	NikkiSerializer(d, func(c *Context) (int, string, gin.H) { return 0, "", gin.H{"a": "b"} })(c)
	assert.NotEmpty(t, r.Result())
}

func TestPlainSerializer(t *testing.T) {
	d := Default("test")
	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)
	c.Request, _ = http.NewRequest(http.MethodGet, "http://0.0.0.0", nil)
	PlainSerializer(d, func(c *Context) (int, string, interface{}) { return 0, "", nil })(c)
	assert.NotEmpty(t, r.Result())
}

func TestTDSSerializer(t *testing.T) {
	d := Default("test")
	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)
	c.Request, _ = http.NewRequest(http.MethodGet, "http://0.0.0.0", nil)
	TDSSerializer(d, func(c *Context) (int, string, gin.H) { return 0, "", gin.H{"a": "b"} })(c)
	assert.NotEmpty(t, r.Result())
}

func TestMethod(t *testing.T) {
	d := Default("test")
	d.Use()
	d.Handle("GET", "/1/")
	d.GET("/2/")
	d.POST("/3/")
	d.PUT("/4/")
	d.DELETE("/5/")
	d.PATCH("/6/")
	d.OPTIONS("/6/")
	d.HEAD("/6/")
}
