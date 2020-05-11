// 这里提供一个cli.go的一个实现，用来方便新项目默认功能
package cli

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewYayagfApp(name, version string, engine *gin.Engine, port int) *App {
	app := NewApp(name, version)

	ServerFactory := func() (*Command, error) {
		c := &Command{
			Run: func(args []string, flags map[string]string) (int, error) {
				if err := engine.Run(fmt.Sprintf(":%v", port)); err != nil {
					return 1, err
				}
				return 0, nil
			},
		}
		return c, nil
	}

	MigrateFactory := func() (*Command, error) {
		c := &Command{
			Run: func(args []string, flags map[string]string) (int, error) {
				return 0, nil
			},
		}
		return c, nil
	}

	app.Commands = map[string]CommandFactory{
		"":        ServerFactory,
		"server":  ServerFactory,
		"migrate": MigrateFactory,
	}

	return app
}
