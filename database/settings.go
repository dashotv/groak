package database

import (
	"encoding/json"
	"errors"

	"github.com/samber/lo"
	"go.etcd.io/bbolt"
)

type Settings struct {
	Downloaders map[string]string `json:"downloaders"`
	Scrapers    []string          `json:"scrapers"`
	Pages       []*Page           `json:"pages"`
}

type Page struct {
	Name       string `json:"name"`
	URL        string `json:"url"`
	Scraper    string `json:"scraper"`
	Downloader string `json:"downloader"`
}

func (s *Settings) AddDownloader(name, url string) {
	if s.Downloaders == nil {
		s.Downloaders = make(map[string]string)
	}
	s.Downloaders[name] = url
}
func (s *Settings) RemoveDownloader(name string) {
	delete(s.Downloaders, name)
}
func (s *Settings) AddScraper(name string) {
	if s.Scrapers == nil {
		s.Scrapers = make([]string, 0)
	}
	if lo.Contains(s.Scrapers, name) {
		return
	}
	s.Scrapers = append(s.Scrapers, name)
}
func (s *Settings) RemoveScraper(name string) {
	s.Scrapers = lo.Filter(s.Scrapers, func(scraper string, i int) bool {
		return scraper != name
	})
}
func (s *Settings) AddPage(page *Page) {
	if s.Pages == nil {
		s.Pages = make([]*Page, 0)
	}
	found := lo.Filter(s.Pages, func(p *Page, i int) bool {
		return p.Name == page.Name
	})
	if len(found) > 0 {
		return
	}
	s.Pages = append(s.Pages, page)
}
func (s *Settings) RemovePage(name string) {
	s.Pages = lo.Filter(s.Pages, func(p *Page, i int) bool {
		return p.Name != name && p.Name != ""
	})
}

func (d *Database) GetSettings() (*Settings, error) {
	settings := &Settings{}
	err := d.client.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("groak"))
		if b == nil {
			return errors.New("bucket not found")
		}
		value := b.Get([]byte("settings"))
		if value == nil {
			return nil
		}
		return json.Unmarshal(value, settings)
	})
	return settings, err
}

func (d *Database) SaveSettings(settings *Settings) error {
	return d.client.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("groak"))
		if err != nil {
			return err
		}
		value, err := json.Marshal(settings)
		if err != nil {
			return err
		}
		return b.Put([]byte("settings"), value)
	})
}
