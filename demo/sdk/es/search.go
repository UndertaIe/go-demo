package es

import (
	"context"
	"fmt"
	"strings"
)

func (ES) SearchByQuery() {
	cli := new(ES).NewClient()
	resp, err := cli.Search(
		cli.Search.WithContext(context.Background()),
		cli.Search.WithIndex(Index),
		cli.Search.WithQuery("new:副总理"),
		cli.Search.WithPretty(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func (ES) SearchByBody() {
	cli := new(ES).NewClient()
	query := `{
		"query": {
			"match": {
				"data": "马斯克"
			}
		}
	}`
	reader := strings.NewReader(query)
	// Perform the search request.
	resp, err := cli.Search(
		cli.Search.WithContext(context.Background()),
		cli.Search.WithIndex(Index),
		cli.Search.WithBody(reader),
		cli.Search.WithPretty(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
