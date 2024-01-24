package cmd

import (
	"fmt"

	"github.com/robfig/cron/v3"

	"github.com/dashotv/groak/scraper"
)

func Cron(p *scraper.Processor, cfg *Config) {
	// run cron
	if cfg.Schedule != "dev" {
		go func() {
			c := cron.New(cron.WithSeconds())
			fmt.Printf("starting cron: %s\n", cfg.Schedule)
			c.AddFunc(cfg.Schedule, func() {
				p.Process()
			})
			c.Start()
		}()
	}

}
