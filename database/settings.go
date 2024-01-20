package database

import (
	"encoding/json"
	"errors"

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

func (d *Database) GetSettings() (*Settings, error) {
	settings := &Settings{}
	err := d.client.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("groak"))
		if b == nil {
			return errors.New("bucket not found")
		}
		value := b.Get([]byte("settings"))
		if value == nil {
			return errors.New("settings not found")
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
