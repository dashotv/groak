/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// scraperCmd represents the scraper command
var scraperCmd = &cobra.Command{
	Use:   "scraper <name>",
	Short: "add scraper",
	Long:  "add scraper",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		defer db.Close()

		settings.AddScraper(name)
		if err := db.SaveSettings(settings); err != nil {
			log.Fatalf("failed to set settings: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(scraperCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scraperCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scraperCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
