package main

import (
	"github.com/UndertaIe/go-demo/demo"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{}

func main() {
	defer func() {
		rootCmd.Execute()
	}()
	demos := map[string]func(){
		"Insert": demo.Desc,
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
