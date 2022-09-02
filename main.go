package main

import (
	"go-demo/dbs"
	"go-demo/demo"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var rootCmd = cobra.Command{}

func main() {
	defer func() {
		rootCmd.Execute()
	}()
	demos := map[string]func(){
		"Insert": demo.Demo,
	}
	AddDemos(demos)
}

func AddDemos(funcs map[string]func() {
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
}
