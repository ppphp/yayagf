package maotai

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestContext(t *testing.T) {
	c := &Context{Context: &gin.Context{}}
	func() {
		defer func() { recover() }()
		c.Error(nil)
	}()
	c.Error(fmt.Errorf("what"))
}
