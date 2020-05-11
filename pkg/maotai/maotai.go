// 利好茅台，奥利给！
package maotai

import (
	"github.com/gin-gonic/gin"
)

type MaoTai struct {
	*gin.Engine
}

func (m *MaoTai) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.Handle(httpMethod, relativePath, handlers...)
}

func (m *MaoTai) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.GET(relativePath, handlers...)
}

func (m *MaoTai) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.POST(relativePath, handlers...)
}

func (m *MaoTai) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.PUT(relativePath, handlers...)
}

func (m *MaoTai) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.DELETE(relativePath, handlers...)
}

func (m *MaoTai) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.PATCH(relativePath, handlers...)
}

func (m *MaoTai) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.OPTIONS(relativePath, handlers...)
}

func (m *MaoTai) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.OPTIONS(relativePath, handlers...)
}

func New() *MaoTai {
	m := &MaoTai{}

	m.Engine = gin.New()
	return m
}

func Default() *MaoTai {
	m := &MaoTai{}
	m.Engine = gin.Default()
	return m
}
