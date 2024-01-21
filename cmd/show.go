/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show settings and config",
	Long:  "show settings and config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Config:\n")
		fmt.Printf("  Schedule:      %s\n", cfg.Schedule)
		fmt.Printf("  Database path: %s\n", cfg.Data)

		fmt.Printf("\nSettings:\n")
		fmt.Printf("  Scrapers:\n")
		for _, s := range settings.Scrapers {
			fmt.Printf("	%-15.15s\n", s)
		}
		fmt.Printf("  Downloaders:\n")
		for n, d := range settings.Downloaders {
			fmt.Printf("	%-15.15s %s\n", n, d)
		}
		fmt.Printf("  Pages:\n")
		for _, p := range settings.Pages {
			fmt.Printf("	%s\n", p.Name)
			fmt.Printf("		URL:        %s\n", p.URL)
			fmt.Printf("		Scraper:    %s\n", p.Scraper)
			fmt.Printf("		Downloader: %s\n", p.Downloader)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
