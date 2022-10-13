package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

var rawJson = []byte(`{
	"name": {"first": "Tom", "last": "Anderson"},
	"age":37,
	"children": ["Sara","Alex","Jack"],
	"friends": [
	  {"first": "James", "last": "Murphy"},
	  {"first": "Roger", "last": "Craig"}
	]
  }`)

func (Parser) JsonParser() {
	doc := gjson.ParseBytes(rawJson)
	nameLast := doc.Get("name.last")
	fmt.Println(nameLast)
	age := doc.Get("age")
	fmt.Println(age)

	chs := doc.Get("children")
	fmt.Println(chs)

	fds := doc.Get("friends")
	fmt.Println(fds)
	fds.ForEach(func(key, item gjson.Result) bool {
		fmt.Println("first: ", item.Get("first"))
		fmt.Println("last: ", item.Get("last"))
		fmt.Println("===")
		return true
	})

	resp, err := http.Get("https://news.cctv.com/2019/07/gaiban/cmsdatainterface/page/china_1.jsonp?cb=china")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	s := string(bytes)
	s2 := string(jsonp2json(bytes))
	if s != "" {
		fmt.Println(s)
		fmt.Println(s2)
	}
	tree := gjson.ParseBytes(jsonp2json(bytes))
	items := tree.Get("data.list")
	items.ForEach(func(key, item gjson.Result) bool {
		fmt.Println(item.Get("title"))
		fmt.Println(item.Get("noknownField"))
		return true
	})
}

func jsonp2json(b []byte) []byte {
	n := len(b)
	var l = 0
	var r = n - 1
	for i := 0; i < n; i++ {
		if b[i] == '{' {
			l = i
			break
		}
	}
	for i := n - 1; i > 0; i-- {
		if b[i] == '}' {
			r = i
			break
		}
	}
	if r < l {
		return b
	}
	return b[l : r+1]
}
