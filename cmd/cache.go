/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cacheCmd represents the cache command
var cacheCmd = &cobra.Command{
	Use:   "cache <bucket>",
	Short: "manage cache data",
	Long:  "manage cache data",
}

func init() {
	rootCmd.AddCommand(cacheCmd)
	cacheCmd.Flags().BoolP("delete", "d", false, "delete from cache")
}
