/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/spf13/cobra"

	"github.com/dashotv/groak/database"
)

var cfg *Config
var db *database.Database
var settings *database.Settings

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "groak",
	Short: "Watch URLs for new videos and download them",
	Long:  "Watch URLs for new videos and download them",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.groak.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var err error
	// Parse environment variables into Config struct
	cfg = &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to parse config: %s\n", err)
	}
	// Open database
	db, err = database.Open(cfg.Data)
	if err != nil {
		log.Fatalf("failed to open db: %s\n", err)
	}
	// Get settings
	settings, err = db.GetSettings()
	if err != nil {
		log.Fatalf("failed to get settings: %s\n", err)
	}
}
