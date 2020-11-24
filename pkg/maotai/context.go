package maotai

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

// (c*Context)Error 记录err的frame
func (c *Context) Error(err error) error {
	if err == nil {
		panic("err is nil")
	}
	_, file, line, _ := runtime.Caller(1)
	err1 := &gin.Error{Err: fmt.Errorf("%v:%v (%w)", file, line, err), Type: gin.ErrorTypePrivate}
	c.Errors = append(c.Errors, err1)
	return err
}
