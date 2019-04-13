package appflag

import (
	"flag"
	"strings"
)

type AppCmd func(args []string) (err error)

type AppFlag struct {
	Name     string
	subFlags map[string]*AppFlag // 记录子命令
	appCmds  map[string]AppCmd   // 记录当前命令
	cmds     map[string]*bool    // 记录当前执行的命令
	flagSet  *flag.FlagSet       // 当前命令解析
}

func NewAppFlag(name string) *AppFlag {
	return &AppFlag{
		Name:     name,
		subFlags: make(map[string]*AppFlag),
		appCmds:  make(map[string]AppCmd),
		cmds:     make(map[string]*bool),
		flagSet:  flag.NewFlagSet(name, flag.ExitOnError),
	}
}

func (this *AppFlag) AddCmd(name, desc string, cmd AppCmd) {
	this.cmds[name] = this.flagSet.Bool(name, false, desc)
	this.appCmds[name] = cmd
}

func (this *AppFlag) Exec(args []string) (err error) {
	if len(args) == 0 {
		this.flagSet.Parse(args)
		this.flagSet.Usage()
		return
	}

	sub := strings.TrimPrefix(args[0], "-")
	if _, ok := this.cmds[sub]; !ok {
		this.flagSet.Parse(args)
		this.flagSet.Usage()
		return
	}

	subFlag := this.subFlags[sub]
	if subFlag != nil {
		// 进入子命令中
		return subFlag.Exec(args[1:])
	}

	// 执行当前命令
	this.flagSet.Parse(args[:1])
	for name, run := range this.cmds {
		if *run {
			cmd := this.appCmds[name]
			return cmd(args[1:])
		}
	}
	return
}

func (this *AppFlag) AddSubFlag(name, desc string, sub *AppFlag) {
	this.subFlags[name] = sub
	this.AddCmd(name, desc, nil) // 添加到当前命令, 以便于提示
}
