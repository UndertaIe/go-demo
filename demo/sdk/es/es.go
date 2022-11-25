package es

import "fmt"

type ES struct{}

// 命令描述
func (ES) Desc() string {
	return "elasticsearch: 全文检索，收集数据、分析数据可视化，监控"
}

// 子命令
func (ES) Name() string {
	return "es"
}

// 文档
func (ES) Docs() {
	overview := "https://www.elastic.co/guide/en/elasticsearch/client/go-api/8.4/overview.html"
	fmt.Println(overview)
}
