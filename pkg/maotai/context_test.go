package maotai

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestContext(t *testing.T) {
	c := &Context{Context: &gin.Context{}}
	func() {
		defer func() { _ = recover() }()
		_ = c.Error(nil)
	}()
	_ = c.Error(fmt.Errorf("what"))
}
