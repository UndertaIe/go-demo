package gorm

import (
	"fmt"

	"github.com/UndertaIe/go-demo/demo/sdk/gorm/model"
)

func (g G) Get() {
	m1 := model.TypeModel{}
	m2 := model.TypeModel{}
	m3 := model.TypeModel{}

	// 获取第一条记录（主键升序）
	db.First(&m1) // 修改后不能再传入修改
	fmt.Println(m1)

	// 获取最后一条记录（主键降序）
	db.Last(&m2)
	fmt.Println(m2)

	// 获取一条记录
	db.Take(&m3)
	fmt.Println(m3)

	m := model.TypeModel{Id: 2}
	db.Find(&m)
	fmt.Println(m)

	m = model.TypeModel{Id: 8}
	db.Find(&m)
	fmt.Println(m)
}

func (g G) List() {
	models := make([]model.TypeModel, 5)
	// 传入切片指针获取全部记录
	ndb := db.Find(&models)

	fmt.Println(ndb.RowsAffected)
	fmt.Println(models)

	m := model.TypeModel{}
	// 传入结构体指针获取单个记录
	db.Find(&m) // equal to db.Take(&m)
	fmt.Println(m)
}

func (g G) Where() {
	m := model.TypeModel{}
	m2 := model.TypeModel{}
	m3 := model.TypeModel{}
	m4 := model.TypeModel{}
	m5 := model.TypeModel{}
	models := make([]model.TypeModel, 5)
	models2 := make([]model.TypeModel, 5)
	models3 := make([]model.TypeModel, 5)
	models4 := make([]model.TypeModel, 5)
	models5 := make([]model.TypeModel, 5)

	// 等值匹配: =
	db.Where("id=?", 1).First(&m)
	fmt.Println(m)

	// 等值匹配: =
	db.Where("str = ?", "").Last(&m2)
	fmt.Println(m2)

	// 等值匹配: <>
	db.Where("f32 <> ?", 0).Last(&m3)
	fmt.Println(m3)

	// 范围查询: IN
	db.Where("id IN ?", []int{2, 3}).Find(&models)
	fmt.Println(m2)

	// 模糊匹配: %
	db.Where("str LIKE ?", "%mi%").Find(&models2)
	fmt.Println(models2)

	// 多条件: AND
	db.Where("f32 = ? AND f64 = ?", 0, 0).Find(&models3)
	fmt.Println(models3)

	// 多条件: OR
	db.Where("f32 = ? OR f64 = ?", 0, 0).Find(&models4)
	fmt.Println(models4)

	// 范围查询：Between
	db.Where("id BETWEEN ? and  ?", 3, 4).Find(&models5)
	fmt.Println(models5)

	// struct条件过滤 tips:当字段为默认值时不会使用该条件
	db.Where(&model.TypeModel{Str: "mikasa", Bo: false}).Find(&m4)
	fmt.Println(`&model.TypeModel{Str: "mikasa"}   | `, m4)

	// map条件过滤 当字段为默认值时使用传入map方式过滤
	db.Where(map[string]interface{}{"str": "mikasa"}).Find(&m5)
	fmt.Println(`map[string]interface{}{"str": "mikasa"}   | `, m5)

}

func (g G) NoWhere() {
	// 主键过滤
	m := model.TypeModel{}
	db.Find(&m, 2)
	fmt.Println(m)

	// Plain SQL
	m2 := model.TypeModel{}
	db.Find(&m2, "str = ?", "mikasa")
	fmt.Println(m2)

	models := make([]model.TypeModel, 5)
	db.Find(&models, "bo = ?", false)
	fmt.Println(models)

	// struct过滤
	m3 := model.TypeModel{}
	db.Find(&m3, &model.TypeModel{Str: "levi"})
	fmt.Println(m3)

	// map过滤（key为mysql表字段）
	m4 := model.TypeModel{}
	db.Find(&m4, map[string]interface{}{"Str": "levi"})
	fmt.Println(m3)
}

// Not()中参数可以为 sql，map，struct, slice
func (g G) Not() {

	// Plain SQL
	m2 := model.TypeModel{}
	db.Not("str = ?", "mikasa").Find(&m2)
	fmt.Println(m2)

	// struct过滤
	m3 := model.TypeModel{}
	db.Not(&model.TypeModel{Str: "levi"}).Find(&m3)
	fmt.Println(m3)

	// map过滤（key为mysql表字段）
	m4 := model.TypeModel{}
	db.Not(map[string]interface{}{"Str": "levi"}).Find(&m4)
	fmt.Println(m4)

	// slice过滤主键记录
	m5 := model.TypeModel{}
	db.Not([]int{2, 3}).Find(&m5)
	fmt.Println(m5)

	fmt.Println("================")

}

// Or()中参数可以为 sql，map，struct, slice
func (g G) Or() {
	m := model.TypeModel{}
	db.Or("id = ?", 2).Find(&m)
	fmt.Println(m)

	models := make([]model.TypeModel, 5)
	db.Or("id = ?", 2).Or("id = ?", 3).Find(&models)
	fmt.Println(models)

	models2 := make([]model.TypeModel, 5)
	db.Or([]int{2, 3, 4}).Find(&models2)
	fmt.Println(models2)
}

// 排序 ： 单字段，多字段
func (g G) Order() {
	models := make([]model.TypeModel, 5)
	db.Order("str").Find(&models) // 排序后 null > 空字符串 > 非空字符串
	fmt.Println(models)
	fmt.Println("models.Str:")
	for _, it := range models {
		if it.Str == "" { // mysql null值转化为空字符串""
			fmt.Print(`""`, " > ")
		} else {
			fmt.Print(it.Str, " > ")
		}

	}
	fmt.Println()

	models2 := make([]model.TypeModel, 5)
	db.Order("str desc, pstr").Find(&models2)
	fmt.Println(models2)
}

func (g G) OffsetLimit() {
	m := model.TypeModel{}
	db.Limit(2).Find(&m)
	fmt.Println(m)

	// Limit(-1) 取消limit限制
	ms := make([]model.TypeModel, 5)
	db = db.Limit(1).Limit(-1).Find(&ms)
	fmt.Println(ms)
	fmt.Println(db.RowsAffected)

	models := make([]model.TypeModel, 4)
	db.Offset(5).Limit(3).Find(&models)
	fmt.Println(models)
	fmt.Println(len(models))

	// Offset(-1) 取消offset限制
	models2 := make([]model.TypeModel, 4)
	db.Offset(5).Offset(-1).Limit(3).Find(&models2)
	fmt.Println(models2)
	fmt.Println(len(models2))

}
