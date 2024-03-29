/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "intialize the database",
	Long:  "intialize the database",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.Setup(); err != nil {
			log.Fatalf("failed to setup db: %s\n", err)
		}

		// p, err := scraper.New(db)
		// if err != nil {
		// 	log.Fatalf("failed to create processor: %s\n", err)
		// }

		// if err := p.Initialize(); err != nil {
		// 	log.Fatalf("failed to initialize processing: %s\n", err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
