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

func New(db *database.Database, settings *database.Settings) (*Processor, error) {
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

func (p *Processor) ProcessSingle(name string) error {
	log.Printf("processing: %s", name)
	for _, page := range p.Settings.Pages {
		if page.Name != name {
			continue
		}

		err := p.processPage(page)
		if err != nil {
			log.Printf("error: %s\n", err)
		}
	}
	log.Printf("processing: complete")
	return nil
}

func (p *Processor) Process() error {
	log.Printf("processing: %d pages", len(p.Settings.Pages))
	for _, page := range p.Settings.Pages {
		if err := p.processPage(page); err != nil {
			log.Printf("error: %s\n", err)
		}
		// log.Printf("finished: %s = %s\n", name, url)
		<-time.After(5 * time.Second)
	}
	log.Printf("processing: complete")
	return nil
}

func (p *Processor) processPage(page *database.Page) error {
	m := NewScraper(page.Scraper)
	if m == nil {
		return fmt.Errorf("scraper not found: %s", page.Scraper)
	}

	// log.Printf("processing: %s = %s\n", page.Name, page.URL)
	for _, v := range m.Read(page.URL) {
		if err := p.Download(page.Name, v, page.Downloader); err != nil {
			log.Printf("error: %s\n", err)
		}
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

	d := NewDownloader(downloader, p.Settings.Downloaders[downloader])
	if d == nil {
		return fmt.Errorf("downloader not found: %s", downloader)
	}

	if !p.init {
		log.Printf("download: %s: %s\n", name, url)
		if err := d.Download(name, url); err != nil {
			_ = p.DB.Set(name, url, fmt.Sprintf("ERROR:%s", err))
			return fmt.Errorf("download: %s", err)
		}
	}

	if err := p.DB.Set(name, url, time.Now().String()); err != nil {
		return fmt.Errorf("db set: %s: %s", url, err)
	}

	return nil
}
