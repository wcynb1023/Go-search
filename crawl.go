package main

import (
	//"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	//"strconv"
)

func init_crawl() *colly.Collector {

	c := colly.NewCollector(
		colly.MaxDepth(4),
		colly.AllowedDomains("wiki.osdev.org"),
	)

	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		//fmt.Println("Visited", r.Request.URL.String())
		vis_list(r.Request.URL.String())
		//fmt.Println("Response:", string(r.Body))
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/") {
			link = e.Request.AbsoluteURL(link)
		}

		// 在这里可以对链接进行进一步处理或过滤

		//fmt.Println("Visiting:", link)

		err := e.Request.Visit(link)
		if err != nil {
			log.Println(err)
		}
		//fmt.Println("Visited:", link)
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		ids := find_key(e.Request.URL.String())

		text := strings.TrimSpace(e.Text)
		if text != "" {
			//fmt.Println(ids , "  <p> :", text)
			text = clean_str(text)
			var documents = []string{text}

			index := BuildInvertedIndex(documents, ids)
			//fmt.Println("Inverted Index:")
			for word, docIDs := range index {
				push_in_db(word, docIDs)
				//fmt.Printf("%s: %v\n", word, docIDs)
			}
		}
	})

	return c
}
