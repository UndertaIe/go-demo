package es

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

const Index = "test_index"

func (ES) Get() {
	cli := new(ES).NewClient()

	resp, err := cli.Get(Index, "5")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.IsError())
	fmt.Println(resp)
}

// http://192.168.3.200:5600/app/dev_tools#/console
func (ES) Create() {
	cli := new(ES).NewClient()
	payloads := []string{`{
			"data": "伸出橄榄枝？乌克兰副总理： 马斯克是最重要的支持者之一"
		}`,
		`{
			"data": "伸出橄榄枝？ 马斯克是最重要的支持者之一"
		}`,
	}

	for idx, payload := range payloads {
		r := strings.NewReader(payload)
		resp, err := cli.Create(Index, cast.ToString(idx), r)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	}
}

func (ES) Update() {
	cli := new(ES).NewClient()
	payload := `{
		"doc": {
			"data": "伸出橄榄枝？乌克兰副总理：马斯克是最重要的支持者之一 update"
		}
	}`
	r := strings.NewReader(payload)
	resp, err := cli.Update(Index, "1", r)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func (ES) CreateForScript() {
	cli := new(ES).NewClient()
	payload := `{
		"rank": 4,
		"data": "for update by script"
	}`
	resp, err := cli.Create(Index, "UpdateByScript", strings.NewReader(payload))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func (ES) UpdateByScript() {
	var updateTmpl = `{
		"script": {
			"source": "if (ctx._source.rank > params.rank ) {ctx._source.rank = params.rank}",
			"params": {"rank": <RANK>}
		}
	}`
	cli := new(ES).NewClient()
	resp, err := cli.Update(
		Index,
		"UpdateByScript",
		// strings.NewReader(strings.ReplaceAll(updateTmpl, "<RANK>", "1")),
		strings.NewReader(strings.ReplaceAll(updateTmpl, "<RANK>", "3")),
		cli.Update.WithPretty(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func (ES) Delete() {
	cli := new(ES).NewClient()

	resp, err := cli.Delete(Index, "UpdateByScript")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
