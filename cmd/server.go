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

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"

	"github.com/dashotv/groak/scraper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run the server",
	Long:  "run the server",
	Run: func(cmd *cobra.Command, args []string) {
		defer db.Close()

		// setup signals
		var stopChan = make(chan os.Signal, 2)
		signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

		p, err := scraper.New(db)
		if err != nil {
			log.Fatalf("failed to create processor: %s\n", err)
		}

		p.Process()

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
	// serverCmd.Flags().BoolP("init", "i", false, "initialize database")
}
