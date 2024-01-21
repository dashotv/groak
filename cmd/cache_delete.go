/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete cache data",
	Long:  "delete cache data",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		bucket := args[0]
		if len(args) == 1 {
			if err := db.DeleteBucket(bucket); err != nil {
				log.Fatalf("failed to delete bucket: %s: %s\n", bucket, err)
			}
			return
		}

		key := args[1]
		if err := db.Delete(bucket, key); err != nil {
			log.Fatalf("failed to delete cache: %s\n", err)
		}
	},
}

func init() {
	cacheCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
