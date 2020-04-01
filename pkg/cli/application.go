// 这里提供一个cli.go的一个实现，用来方便新项目默认功能
package cli

var ApplicationCommands = &Command{
	Commands: map[string]CommandFactory{
		"":        ServerFactory,
		"server":  ServerFactory,
		"migrate": MigrateFactory,
	},
}

func ServerFactory() (*Command, error) {
	c := &Command{
		Run: func(args []string) (int, error) {
			return 0, nil
		},
	}
	return c, nil
}

func MigrateFactory() (*Command, error) {
	c := &Command{
		Run: func(args []string) (int, error) {
			return 0, nil
		},
	}
	return c, nil
}
