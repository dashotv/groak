/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/dashotv/groak/database"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page <name> <url>",
	Short: "add page",
	Long:  "add page",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		u := args[1]

		if name == "" || u == "" {
			log.Fatal("name and url are required")
		}

		if _, err := url.Parse(u); err != nil {
			log.Fatalf("invalid url: %s\n", err)
		}

		defer db.Close()

		scraper, err := cmd.Flags().GetString("scraper")
		if err != nil {
			log.Fatalf("failed to get scraper flag: %s\n", err)
		}

		downloader, err := cmd.Flags().GetString("downloader")
		if err != nil {
			log.Fatalf("failed to get downloader flag: %s\n", err)
		}

		settings.AddPage(&database.Page{
			Name:       name,
			URL:        u,
			Scraper:    scraper,
			Downloader: downloader,
		})

		if err := db.SaveSettings(settings); err != nil {
			log.Fatalf("failed to set settings: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pageCmd.Flags().StringP("scraper", "s", "myanime", "set the scraper to use for this page")
	pageCmd.Flags().StringP("downloader", "d", "metube", "set the downloader to use for this page")
}
