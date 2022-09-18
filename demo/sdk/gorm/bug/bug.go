package bug

import (
	"fmt"

	"github.com/UndertaIe/go-demo/db/mysql"
	"github.com/UndertaIe/go-demo/demo/sdk/gorm/model"
)

var db = mysql.Db

type B struct{}

func (b B) Name() string {
	return "bug"
}

func (b B) Desc() string {
	return "gorm fix"
}

// 	bug: error 1406:data too long for column 'salt' at row 1
// 其中 salt 在mysql中为 char(4) 类型
// 最终结论是make([]byte, 4)中的4是length，即切片中前四位为编码为0，0，0，0生成salt值时即使print的是4位字符串，但其实是8个字符，往数据库写时也是长度为8，导致写入失败
// make([]byte,n)中n为切片length
// make([]byte,x,y)中x为length，y为cap
func InsertChar() {
	m := model.CharTest{C: "xxxx"}
	db.Model(m).Create(&m)
	x := make([]byte, 4)
	y := make([]byte, 5, 10)
	fmt.Printf("len(x)=%d, cap(x)=%d\n", len(x), cap(x))
	fmt.Printf("len(y)=%d, cap(y)=%d\n", len(y), cap(y))
}
