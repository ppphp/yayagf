// 别人的subcommand功能都用不爽，手写个傻乎乎的cli。。主要是想把subcommand写到command本身的map里
// TODO: cli flags support 关键是我也没用--help这种功能。。。
// 实现一个Command的interface，然后手写Command实现，继承一个自带map的subcommands来
package cli

type CommandFactory func() (*Command, error)

type Command struct {
	// 非常有用，关键中的关键
	Commands map[string]CommandFactory
	// 普通的args
	Args []string
	// 运行函数，mute when subcommand exists
	Run func(args []string) (int, error)
}

// whole lifetime for a command, preserve for hooks or something else
func (c *Command) exec() (int, error) {
	if c.Commands == nil {
		return c.Run(c.Args)
	}
	if len(c.Args) == 0 {
		if s, ok := c.Commands[""]; !ok {
			if c.Run!= nil {
				return c.Run(c.Args)
			} else {
				return 1, nil
			}
		} else {
			if f, err := s(); err != nil {
				return 1, err
			} else {
				f.Args = c.Args
				return f.exec()
			}
		}
	} else {
		if s, ok := c.Commands[c.Args[0]]; !ok {
			// preserve for help function
			return 1, nil
		} else {
			if f, err := s(); err != nil {
				return 1, err
			} else {
				if len(c.Args) > 0 {
					f.Args = c.Args[1:]
				}
				return f.exec()
			}
		}
	}
}

// 根command，当然也可以用来做普通command，就是个例子
type App struct {
	Name    string
	Version string
	*Command
}

func NewApp(name, version string) *App {
	a := &App{Name: name, Version: version, Command: &Command{}}
	return a
}

func (a *App) Run() (int, error) {
	return a.exec()
}
