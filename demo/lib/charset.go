package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
)

// comment: 修改html编码为utf8
func (Lib) CharSetDemo() {
	resp, _ := http.Get("http://www.people.com.cn/GB/59476/index.html")
	bytes, _ := ioutil.ReadAll(resp.Body)
	s, n, ok := charset.DetermineEncoding(bytes, "utf8")
	fmt.Println(s, n, ok)
}
