/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// downloaderCmd represents the downloader command
var downloaderCmd = &cobra.Command{
	Use:   "downloader <name> <url>",
	Short: "add downloader",
	Long:  "add downloader",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		url := args[1]

		defer db.Close()

		settings.AddDownloader(name, url)

		if err := db.SaveSettings(settings); err != nil {
			log.Fatalf("failed to set settings: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloaderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloaderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloaderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
