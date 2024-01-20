package myanime

import (
	"fmt"

	"github.com/gocolly/colly"
)

func New(url string) *MyAnime {
	return &MyAnime{
		url: url,
		col: colly.NewCollector(),
	}
}

type MyAnime struct {
	url string
	col *colly.Collector
}

func (m *MyAnime) Read() []string {
	urls := []string{}
	m.col.OnHTML("article", func(e *colly.HTMLElement) {
		urls = append(urls, e.ChildAttr("a", "href"))
	})
	m.col.OnError(func(r *colly.Response, err error) {
		fmt.Printf("error: %s\n", err)
	})
	m.col.Visit(m.url)
	return urls
}
