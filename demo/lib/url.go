package lib

import (
	"net/url"
)

func UrlJoin(baseUrl string, joinUrl string) (string, error) {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return "", nil
	}
	join, err := url.Parse(joinUrl)
	if err != nil {
		return "", nil
	}
	return base.ResolveReference(join).String(), nil
}

func (Lib) UrlJoinTest() {
	baseUrl := "https://www.chinanews.com.cn/scroll-news/news1.html"
	joinUrl := "/gn/2022/10-13/9872426.shtml"
	result := "https://www.chinanews.com.cn/gn/2022/10-13/9872426.shtml"
	if r, _ := UrlJoin(baseUrl, joinUrl); r != result {
		panic("UrlJoin error")
	}
}

func (Lib) UrlJoinTest2() {
	baseUrl := "https://www.chinanews.com.cn/scroll-news/news1.html"
	joinUrl := "https://www.chinanews.com.cn/gn/2022/10-13/9872426.shtml"
	result := "https://www.chinanews.com.cn/gn/2022/10-13/9872426.shtml"
	if r, _ := UrlJoin(baseUrl, joinUrl); r != result {
		panic("UrlJoin error")
	}
}
