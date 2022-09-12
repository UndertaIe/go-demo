package main

import (
	"fmt"
	"go/importer"
	"go/types"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{}

func main() {
	defer func() {
		rootCmd.Execute()
	}()
	pkg, err := importer.Default().Import("time")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		if name == "Time" {
			obj := scope.Lookup(name)
			if tn, ok := obj.Type().(*types.Named); ok {
				for i := 0; i < tn.NumMethods(); i++ {
					fmt.Println()
					m := tn.Method(i)
					m.Pkg()
					m.FullName()
					m.Exported()
				}
				fmt.Printf("%#v\n", tn.NumMethods())
			}
		}
	}
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
