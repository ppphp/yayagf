package yac

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	eventID string
}

func FromGin(c *gin.Context) *Context {
	ctx := &Context{Context: c}
	return ctx
}
