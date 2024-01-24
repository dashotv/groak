/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/dashotv/groak/scraper"
)

type kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

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

		go func() {
			p.Process()
		}()

		Cron(p, cfg)
		Router(p, db, settings)

		for {
			select {
			case <-stopChan:
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
