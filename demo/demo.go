package demo

import (
	"bytes"
	"reflect"

	"github.com/spf13/cobra"
)

var Rdemo = new(cobra.Command)

type demo interface {
	Desc() string
	Name() string
}

func Fire(rdemo *cobra.Command, d demo) *cobra.Command {
	v := reflect.ValueOf(d)
	t := reflect.TypeOf(d)

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		Use:   d.Name(),
		Short: d.Desc(),
		Long:  d.Desc(),
	}
	for i := 0; i < v.NumMethod(); i++ {
		m := t.Method(i)
		b := bytes.NewBufferString(m.PkgPath)
		// b.WriteString(":")
		b.WriteString(m.Name)
		vv := v.Method(i) //回调函数要尽量避免使用外部的参数
		cmd.AddCommand(&cobra.Command{
			Run: func(cmd *cobra.Command, args []string) {
				// v.Method(i).Call(make([]reflect.Value, 0)) // error: 这种写法会导致回调函数会依赖外部的i参数，当真正执行时会i跟之前遍历时不一样
				vv.Call(make([]reflect.Value, 0))
			},
			Use:   m.Name,
			Short: b.String(),
			Long:  b.String(),
		})
	}
	rdemo.AddCommand(cmd)
	return cmd
}
