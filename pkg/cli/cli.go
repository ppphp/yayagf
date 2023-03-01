// 别人的subcommand功能都用不爽，手写个傻乎乎的cli。。主要是想把subcommand写到command本身的map里
// 实现一个Command的interface，然后手写Command实现，继承一个自带map的subcommands来
package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/ppphp/yayagf/pkg/meta"
)

type CommandFactory func() (*Command, error)

type Command struct {
	// 注册的入口
	Commands map[string]CommandFactory
	// 输入的参数，字符串数组
	RawArgs []string
	// 运行函数，mute when subcommand exists
	Run func(args []string, flags map[string]string) (int, error)

	// 单步解析的参数
	// 当前的操作
	oneStepCommand string
	// 操作之前的flags
	oneStepFlags map[string]string
	// 当前剩下的参数
	oneStepOtherArgs []string

	// 剩下一起解析的参数
	allCommands []string
	// 操作之前的flags
	allFlags map[string]string
}

// whole lifetime for a command, preserve for hooks or something else
func (c *Command) exec(cargs []string) (int, error) {
	c.parseArgs(cargs)

	if c.Commands == nil {
		if c.Run != nil {
			return c.Run(c.allCommands, c.allFlags)
		}
		return 1, nil
	}

	if len(c.RawArgs) == 0 {
		s, ok := c.Commands[""]
		if !ok {
			if c.Run != nil {
				return c.Run(c.allCommands, c.allFlags)
			}
			return 1, nil
		}

		f, err := s()

		if err != nil {
			return 1, err
		}

		return f.exec(c.RawArgs[1:])
	}

	s, ok := c.Commands[c.oneStepCommand]

	if !ok {
		// preserve for help function
		return 1, nil
	}

	f, err := s()

	if err != nil {
		return 1, err
	}

	return f.exec(c.oneStepOtherArgs)
}

// with root command
func (c *Command) parseArgs(args []string) {
	c.RawArgs = args
	c.oneStepFlags = map[string]string{}
	c.oneStepOtherArgs = []string{}

	foundCommand := false
	for i := range args {
		if foundCommand {
			c.oneStepOtherArgs = append(c.oneStepOtherArgs, args[i])
		} else {
			if !strings.HasPrefix(args[i], "-") {
				c.oneStepCommand = args[i]
				foundCommand = true
				continue
			}

			f := strings.SplitN(strings.TrimLeft(args[i], "-"), "=", 2)

			if len(f) == 1 {
				c.oneStepFlags[f[0]] = ""
			} else {
				c.oneStepFlags[f[0]] = f[1]
			}
		}
	}

	c.allFlags = map[string]string{}
	c.allCommands = []string{}

	for i := range args {
		if strings.HasPrefix(args[i], "-") {
			f := strings.SplitN(strings.TrimLeft(args[i], "-"), "=", 2)

			if len(f) == 1 {
				c.allFlags[f[0]] = ""
			} else {
				c.allFlags[f[0]] = f[1]
			}
		} else {
			c.allCommands = append(c.allCommands, args[i])
		}
	}
}

// 根command，加入一些二进制相关的帮助函数
type App struct {
	Name string
	*Command
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
	md := meta.Get()
	return fmt.Sprintf("%v %v, digested %v built by %v %v on %v %v at %v with intranet %v", a.Name, md.Version, md.MD5, md.GoCompiler, md.GoVersion, md.GoOS, md.GoArch, md.BuildAt, md.Local)
}
