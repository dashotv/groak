package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func NewMyAnime() *MyAnime {
	return &MyAnime{
		col: colly.NewCollector(),
	}
}

type MyAnime struct {
	col *colly.Collector
}

func (m *MyAnime) Read(url string) []string {
	urls := []string{}
	m.col.OnHTML("article", func(e *colly.HTMLElement) {
		urls = append(urls, e.ChildAttr("a", "href"))
	})
	m.col.OnError(func(r *colly.Response, err error) {
		fmt.Printf("error: %s\n", err)
	})
	m.col.Visit(url)
	return urls
}
