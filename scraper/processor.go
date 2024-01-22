package scraper

import (
	"fmt"
	"log"
	"time"

	"github.com/dashotv/groak/database"
)

type Processor struct {
	DB       *database.Database
	Settings *database.Settings
	init     bool
}

type Scraper interface {
	Read(url string) []string
}

type Downloader interface {
	Download(name, url string) error
}

func New(db *database.Database) (*Processor, error) {
	settings, err := db.GetSettings()
	if err != nil {
		return nil, err
	}

	p := &Processor{
		DB:       db,
		Settings: settings,
	}

	return p, nil
}

func NewScraper(name string) Scraper {
	switch name {
	case "myanime":
		return NewMyAnime()
	}
	return nil
}

func NewDownloader(name, url string) Downloader {
	switch name {
	case "metube":
		return NewMetube(url)
	}
	return nil
}

func (p *Processor) Initialize() error {
	p.init = true
	return p.Process()
}

func (p *Processor) Process() error {
	log.Printf("processing...")
	for _, page := range p.Settings.Pages {
		m := NewScraper(page.Scraper)
		if m == nil {
			return fmt.Errorf("scraper not found: %s", page.Scraper)
		}

		log.Printf("processing: %s = %s\n", page.Name, page.URL)
		for _, v := range m.Read(page.URL) {
			if err := p.Download(page.Name, v, page.Downloader); err != nil {
				log.Printf("error: %s\n", err)
			}
		}

		// log.Printf("finished: %s = %s\n", name, url)
		<-time.After(5 * time.Second)
	}
	return nil
}

func (p *Processor) Download(name, url, downloader string) error {
	val, err := p.DB.Get(name, url)
	if err != nil {
		log.Printf("error: db get: %s: %s\n", url, err)
		return fmt.Errorf("error: db get: %s: %s\n", url, err)
	}
	if val != "" {
		// log.Printf("skipping: %s: %s\n", v, val)
		return nil
	}

	d := NewDownloader(downloader, url)
	if d == nil {
		return fmt.Errorf("downloader not found: %s", downloader)
	}

	if p.init {
		log.Printf("init: %s: %s\n", name, url)
	} else {
		log.Printf("download: %s: %s\n", name, url)
		if err := d.Download(name, url); err != nil {
			_ = p.DB.Set(name, url, err.Error())
			return fmt.Errorf("download: %s", err)
		}
	}

	if err := p.DB.Set(name, url, time.Now().String()); err != nil {
		return fmt.Errorf("db set: %s: %s", url, err)
	}

	return nil
}
