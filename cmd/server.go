/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v10"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"

	"github.com/dashotv/groak/database"
	"github.com/dashotv/groak/scraper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run the server",
	Long:  "run the server",
	Run: func(cmd *cobra.Command, args []string) {
		// Parse environment variables into Config struct
		cfg := &Config{}
		if err := env.Parse(cfg); err != nil {
			fmt.Printf("failed to parse config: %s\n", err)
			os.Exit(1)
		}

		// setup signals
		var stopChan = make(chan os.Signal, 2)
		signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

		db, err := database.Open(cfg.Data)
		if err != nil {
			log.Fatalf("failed to open db: %s\n", err)
		}
		defer db.Close()

		p := &scraper.Processor{
			DB: db,
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
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
