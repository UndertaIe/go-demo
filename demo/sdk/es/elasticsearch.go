package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	Addrs = []string{"http://192.168.3.200:9200"}
	// note: /usr/local/es/elasticsearch-8.4.3/bin/
	// note: ./elasticsearch-service-tokens  create 'elastic/enterprise-search-server' 'go-client' 生成 service token
	ServiceToken = "AAEAAWVsYXN0aWMvZW50ZXJwcmlzZS1zZWFyY2gtc2VydmVyL2dvLWNsaWVudDo5aTVtODVvWVJrLUtLQWdLWUd0UHNn"
)

func (ES) NewClient() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: Addrs,
		// ServiceToken: ServiceToken,
	}
	cli, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return cli
}

func (ES) ReadmeCode() {
	log.SetFlags(0)
	var (
		r map[string]interface{}
	)
	e := ES{}
	cli := e.NewClient()
	res, err := cli.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	var wg sync.WaitGroup
	for i, title := range []string{"Test 中国", "Test 中国人民"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Build the request body.
			data, err := json.Marshal(struct{ Title string }{Title: title})
			if err != nil {
				log.Fatalf("Error marshaling document: %s", err)
			}
			fmt.Println(string(data))
			// Set up the request object.
			req := esapi.IndexRequest{
				Index: Index,
				// DocumentID: strconv.Itoa(i + 1),
				Body:    bytes.NewReader(data),
				Refresh: "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), cli)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
					log.Println(r)
				}
			}
		}(i, title)
	}
	wg.Wait()

	log.Println(strings.Repeat("-", 37))

}

func (ES) SearchDemo() {
	var (
		r map[string]interface{}
	)
	cfg := elasticsearch.Config{
		Addresses: Addrs,
		// ServiceToken: ServiceToken,
	}
	cli, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	query := `{
		"query": {
			"match": {
				"Title": "中国"
			}
		}
	}`
	reader := strings.NewReader(query)
	// Perform the search request.
	res, err := cli.Search(
		cli.Search.WithContext(context.Background()),
		cli.Search.WithIndex(Index),
		cli.Search.WithBody(reader),
		cli.Search.WithTrackTotalHits(true),
		cli.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	log.Println(r)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))

}
