package gorm

import "github.com/UndertaIe/go-demo/db/mysql"

var db = mysql.Db

type G struct{}

func (g G) Name() string {
	return "gorm"
}

func (g G) Desc() string {
	return "gorm 示例"
}
