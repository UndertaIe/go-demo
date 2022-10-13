package parser

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func (Parser) HtmlParser() {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	goquery.NewDocument("http://metalsucks.net")
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".left-content article").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find(".post-title a").Text()
		link, _ := s.Find(".post-title a").Attr("href")
		publishTime := s.Find("time").Text()
		abs := s.Find(".excerpt p").Text()
		fmt.Printf("Review %d: \n", i)
		blank := "    "
		fmt.Println(blank, title)
		fmt.Println(blank, link)
		fmt.Println(blank, publishTime)
		fmt.Println(blank, abs)
	})
}
