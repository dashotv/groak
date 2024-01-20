package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/dashotv/groak/myanime"
)

type Processor struct {
	cfg *Config
	db  *Database
}

func (p *Processor) Process() {
	for name, url := range p.cfg.URLs {
		// log.Printf("processing: %s = %s\n", name, url)
		m := myanime.New("https://" + url)
		for _, v := range m.Read() {
			val, err := p.db.Get(name, v)
			if err != nil {
				log.Printf("error: db get: %s: %s\n", v, err)
				continue
			}
			if val != "" {
				// log.Printf("skipping: %s: %s\n", v, val)
				continue
			}

			if !p.cfg.Initialize {
				// skip downloads on first run
				log.Printf("download: %s: %s\n", name, v)
				if err := p.Download(name, v); err != nil {
					log.Printf("error: %s", err)
					continue
				}
			}

			if err := p.db.Set(name, v, time.Now().String()); err != nil {
				log.Printf("error: db set: %s: %s", v, err)
				continue
			}
		}

		// log.Printf("finished: %s = %s\n", name, url)
		<-time.After(5 * time.Second)
	}
}

type Download struct {
	URL       string `json:"url"`
	AutoStart bool   `json:"auto_start"`
	Quality   string `json:"quality"`
	Format    string `json:"format"`
	Name      string `json:"custom_name_prefix"`
}

type Response struct {
	Status string `json:"status"`
}

func (p *Processor) Download(name, url string) error {
	client := resty.New()
	resp, err := client.R().
		SetBody(&Download{url, false, "best", "any", name}).
		Post(p.cfg.Metube)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("request failed: %d: %s: %s", resp.StatusCode(), resp.Status(), resp.String())
	}
	return nil
}
