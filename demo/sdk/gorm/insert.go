package gorm

import (
	"github.com/UndertaIe/go-demo/demo/sdk/gorm/model"
)

// 单行插入
func (g G) Insert() {
	m := model.TypeModel{}
	db.Create(&m)
}

// 批插入
func (g G) InsertBatch() {
	models := make([]model.TypeModel, 3)
	db.CreateInBatches(&models, 3)
}

// 选择model字段插入
func (g G) SelectInsert() {
	m := model.TypeModel{ModelId: 3, F32: 1.23}
	db.Select("ModelId").Create(&m)
}

// 忽略model字段插入
func (g G) OmitInsert() {
	m := model.TypeModel{ModelId: 3, F32: 1.23}
	db.Omit("ModelId").Create(&m)
}
