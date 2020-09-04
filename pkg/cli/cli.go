// 别人的subcommand功能都用不爽，手写个傻乎乎的cli。。主要是想把subcommand写到command本身的map里
// 实现一个Command的interface，然后手写Command实现，继承一个自带map的subcommands来
package cli

import (
	"fmt"
	"gitlab.papegames.com/fengche/yayagf/pkg/meta"
	"os"
	"strings"
)

type CommandFactory func() (*Command, error)

type Command struct {
	// 非常有用，关键中的关键
	Commands map[string]CommandFactory
	// 普通的args
	Args []string
	// 普通的flags
	Flags map[string]string
	// 运行函数，mute when subcommand exists
	Run func(args []string, flags map[string]string) (int, error)
}

// whole lifetime for a command, preserve for hooks or something else
func (c *Command) exec(cargs []string) (int, error) {
	c.parseArgs(cargs)

	if c.Commands == nil {
		return c.Run(c.Args, c.Flags)
	}

	if len(c.Args) == 0 {
		s, ok := c.Commands[""]
		if !ok {
			if c.Run != nil {
				return c.Run(c.Args, c.Flags)
			}

			return 1, nil
		}

		f, err := s()

		if err != nil {
			return 1, err
		}

		return f.exec(c.Args[1:])
	}

	s, ok := c.Commands[c.Args[0]]

	if !ok {
		// preserve for help function
		return 1, nil
	}

	f, err := s()

	if err != nil {
		return 1, err
	}

	return f.exec(c.Args[1:])
}

func (c *Command) parseArgs(args []string) {
	c.Flags = map[string]string{}

	for i := range args {
		if !strings.HasPrefix(args[i], "-") {
			c.Args = append([]string{}, args[i:]...)
			return
		}

		f := strings.SplitN(strings.TrimPrefix(args[i], "-"), "=", 2)

		if len(f) == 1 {
			c.Flags[f[0]] = ""
		} else {
			c.Flags[f[0]] = f[1]
		}
	}
}

// 根command，当然也可以用来做普通command，就是个例子
type App struct {
	Name string
	Meta meta.Meta
	*Command
}

func NewApp(name, version string) *App {
	a := &App{Name: name, Meta: meta.Get(), Command: &Command{}}
	return a
}

func (a *App) Run() (int, error) {
	if os.Getenv("Meta") != "mute" {
		fmt.Println(a.PrintMeta())
	}
	return a.RunArgs(os.Args[1:])
}

func (a *App) RunArgs(args []string) (int, error) {
	return a.exec(args)
}

func (a *App) PrintMeta() string {
	return fmt.Sprintf("%v %v, digested %v built by %v %v on %v %v at %v with intranet %v", a.Name, a.Meta.Version, a.Meta.MD5, a.Meta.GoCompiler, a.Meta.GoVersion, a.Meta.GoOS, a.Meta.GoArch, a.Meta.BuildAt, a.Meta.Local)
}
