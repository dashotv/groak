package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v10"
	"github.com/dotenv-org/godotenvvault"
	"github.com/robfig/cron/v3"
)

func main() {
	// Load .env file and populate environment variables
	if err := godotenvvault.Load(); err != nil {
		fmt.Printf("failed to load env: %s\n", err)
		os.Exit(1)
	}

	// Parse environment variables into Config struct
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("failed to parse config: %s\n", err)
		os.Exit(1)
	}

	// setup signals
	var stopChan = make(chan os.Signal, 2)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	db, err := NewDatabase(cfg.Data)
	if err != nil {
		log.Fatalf("failed to open db: %s\n", err)
	}
	defer db.Close()

	p := &Processor{
		cfg: cfg,
		db:  db,
	}

	initialized, err := db.Initialized()
	if err != nil {
		log.Fatalf("failed to check initialized: %s\n", err)
	}
	if !initialized {
		log.Printf("initializing\n")
		cfg.Initialize = true
	}

	// first run
	p.Process()

	if cfg.Initialize {
		log.Printf("initialized\n")
		if err := db.SetInitialized(); err != nil {
			log.Fatalf("failed to set initialized: %s\n", err)
		}
	}

	// run cron
	if cfg.Schedule == "dev" {
		return
	}

	c := cron.New(cron.WithSeconds())
	fmt.Printf("starting cron: %s\n", cfg.Schedule)
	c.AddFunc(cfg.Schedule, func() {
		p.Process()
	})
	c.Start()

	for {
		select {
		case <-stopChan:
			c.Stop()
			println("stopped")
			return
		}
	}
}
