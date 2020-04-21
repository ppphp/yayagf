// 一个swagger的builder
// 分三个部分，handler，route和基本信息
// 基本信息就抄swagger
// route应该是读程序里的注册route
// handler是在函数上的doc
// 最后build的时候合成一个
package builder

import "github.com/go-openapi/spec"

type Route struct {

}

type Builder struct {
	swagger *spec.Swagger
}

func (b *Builder) AddRoute(route, method, accept, income, outcome string) {

}

func New(swagger *spec.Swagger) *Builder {
	b := &Builder{
		swagger: swagger,
	}
	return b
}
