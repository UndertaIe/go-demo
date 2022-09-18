package gorm

import (
	"fmt"

	"github.com/UndertaIe/go-demo/demo/sdk/gorm/model"
)

func (g G) Group() {
	m := model.GroupModel{}
	r := make([]model.GroupResult, 5)
	// .Model(&m)指明操作的表名
	db.Model(&m).Group("str").Select(`str, count(num) as total`).Find(&r)
	fmt.Println(r)

	r2 := make([]model.GroupResult, 5)
	// .Model(&m)指明操作的表名
	db.Model(&m).Group("str").Select(`str, avg(num) as "avg"`).Find(&r2)
	fmt.Println(r2)
}

func (g G) Distinct() {
	var arr []string
	db.Model(&model.GroupModel{}).Distinct("str").Scan(&arr)
	fmt.Println(arr)
}
