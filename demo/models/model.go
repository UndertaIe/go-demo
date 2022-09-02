package models

import "time"

type TypeModel struct {
	Id      int
	ModelId int
	F32     float32
	F64     float64
	Str     string
	Pstr    *string
	Bo      bool
	T1      *time.Time
	T2      *time.Time
	T3      *time.Time
}

func (t TypeModel) TableName() string { // orm framework will get ralational database table name by calling this method
	return "type_model"
}

type GroupModel struct {
	ID  int
	Str string
	Num int
}

type GroupResult struct {
	Str   string
	Total int
	Avg   float32
}

func (t GroupModel) TableName() string { // orm framework will get ralational database table name by calling this method
	return "group_model"
}

type JoinModel struct {
	ID  int
	Str string
	Jid int
}

func (t JoinModel) TableName() string {
	return "join_model"
}

type JoinModel2 struct {
	ID   int
	Str2 string
}

func (t JoinModel2) TableName() string {
	return "join_model2"
}

type JoinResult struct {
	Id   int
	Jid  int
	Str  string
	Str2 string
}

type CharTest struct {
	C string `json:"c" gorm:"column:c"`
}

func (t CharTest) TableName() string {
	return "char_test"
}
