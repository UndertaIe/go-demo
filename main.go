package main

import (
	"go-demo/dbs"
	"go-demo/demo"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{}

func main() {
	defer func() {
		rootCmd.Execute()
	}()
	demos := map[string]func(){
		"Insert":        demo.Demo,
		"Md5Demo":       demo.Md5Demo,
		"TestSliceJoin": demo.TestSliceJoin,
		"FuncType":      demo.FuncType,
		"TestRedis":     demo.TestRedis,
	}
	AddDemos(demos)
}

func AddDemos(funcs map[string]func()) {
	for fn, f := range funcs {
		AddDemo(fn, f)
	}
}

func AddDemo(fn string, f func()) {
	rootCmd.AddCommand(&cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			f()
		},
		Use: fn,
	})
}

func init() {
	_db, err := dbs.NewDBEngine()
	if err != nil {
		panic(err)
	}
	dbs.Db = _db
	dbs.Conn, err = dbs.NewRedisConn()
	if err != nil {
		panic(err)
	}
}
