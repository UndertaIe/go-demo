package gorm

import (
	"fmt"

	"github.com/UndertaIe/go-demo/demo/sdk/gorm/model"
)

func (g G) Join() {
	m := model.JoinModel{}
	rows, err := db.Model(&m).Joins(
		"join join_model2 on join_model.jid = join_model2.id").Select(
		"join_model.id, join_model.str, join_model.jid, join_model2.str2").Rows()
	if err != nil {
		panic(err)
	}
	var results []model.JoinResult // 创建一个切片
	for rows.Next() {
		result := model.JoinResult{}
		err = rows.Scan(&result.Id, &result.Str, &result.Jid, &result.Str2)
		if err != nil {
			panic(err)
		}
		results = append(results, result)
	}
	fmt.Println(results)
}
