package protocolbuffer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

type PB struct{}

func (b PB) Desc() string {
	return "统一数据交换格式，实现高效率的跨平台，跨语言数据通信"
}

func (b PB) Name() string {
	return "protocolbuffer"
}

func (b PB) Docs() {
	tutorials := "https://developers.google.com/protocol-buffers/docs/gotutorial"
	fmt.Println(tutorials)
}

func (b PB) Demo() {
	p := Person{
		Name:  "mikasa",
		Id:    1,
		Email: "xxx@email.com",
	}
	fmt.Println("proto.marshal Person: ", p)
	bytes, err := proto.Marshal(&p)
	if err != nil {
		panic(err)
	}
	var fn = "./data/pb.dat"
	err = ioutil.WriteFile(fn, bytes, 0644)
	if err != nil {
		panic(err)
	}
	bytes2, _ := ioutil.ReadFile(fn)

	var p2 Person
	proto.Unmarshal(bytes2, &p2)
	fmt.Println("proto.Unmarshal result: ", p2)
}
