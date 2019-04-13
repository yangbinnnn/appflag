package appflag

import (
	"flag"
	"fmt"
	"testing"
	"time"
)

func TestAppFlag(t *testing.T) {
	app := NewAppFlag("app")
	app.AddCmd("do", "just do it", func(args []string) (err error) {
		fmt.Println("just do it", args)
		return
	})

	subHello := NewAppFlag("hello")
	subHello.AddCmd("say", "say something", func(args []string) (err error) {
		var msg string
		param := flag.NewFlagSet("say", flag.ExitOnError)
		param.StringVar(&msg, "msg", "hello world", "msg your want to say") // use -h for help
		param.Parse(args)
		fmt.Println("say: ", msg)
		return
	})
	subHello.AddCmd("hi", "hi to name", func(args []string) (err error) {
		var name string
		param := flag.NewFlagSet("hi", flag.ExitOnError)
		param.StringVar(&name, "name", "", "give me an name!")
		param.Parse(args)
		if len(args) == 0 { // show help
			param.Usage()
			return
		}

		fmt.Printf("hi: %s\n", name)
		return
	})

	migrate := NewAppFlag("migrate")
	migrate.AddCmd("mysql", "migrate mysql", func(args []string) (err error) {
		fmt.Println("migrate mysql...")
		return
	})
	migrate.AddCmd("mongo", "migrate mongo", func(args []string) (err error) {
		fmt.Println("migrate mongo...")
		return
	})

	deepcmd := NewAppFlag("deepcmd")
	deepcmd.AddCmd("sleep", "go to sleep", func(args []string) (err error) {
		var sec int64
		param := flag.NewFlagSet("sleep", flag.ExitOnError)
		param.Int64Var(&sec, "second", int64(0), "sleep second")
		param.Parse(args)
		if len(args) == 0 {
			param.Usage()
			return
		}

		fmt.Println("sleep...")
		time.Sleep(time.Duration(sec) * time.Second)
		fmt.Println("sleep done.")
		return
	})

	migrate.AddSubFlag("deepcmd", "deep cmd", deepcmd)

	app.AddSubFlag(subHello.Name, "hello cmd set", subHello)
	app.AddSubFlag(migrate.Name, "migrate cmd set", migrate)
	app.Exec([]string{"-migrate", "-deepcmd", "-sleep", "-second", "1"})
}
