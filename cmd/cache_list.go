/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list cache data",
	Long:  "list cache data",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if all, err := cmd.Flags().GetBool("all"); err == nil && all {
			eachBucket(func(bucket string) {
				log.Printf("-  %s", bucket)
				eachItem(bucket, func(key string) {
					val, err := db.Get(bucket, key)
					if err != nil {
						log.Printf("   -  %s failed to get value: %s", key, err)
						return
					}
					log.Printf("   -  %s %s", key, val)
				})
			})
			return
		}

		if len(args) == 0 {
			eachBucket(func(bucket string) {
				log.Printf("-  %s", bucket)
			})
		}

		bucket := args[0]
		eachItem(bucket, func(key string) {
			val, err := db.Get(bucket, key)
			if err != nil {
				log.Printf("   -  %s failed to get value: %s", key, err)
				return
			}
			log.Printf("   -  %s %s", key, val)
		})
	},
}

func eachBucket(fn func(string)) {
	buckets, err := db.ListBuckets()
	if err != nil {
		fmt.Printf("failed to list buckets: %s\n", err)
		return
	}

	for _, bucket := range buckets {
		if bucket == "groak" {
			continue
		}
		fn(bucket)
	}
}

func eachItem(bucket string, fn func(string)) {
	keys, err := db.List(bucket)
	if err != nil {
		fmt.Printf("failed to list keys: %s\n", err)
		return
	}

	for _, key := range keys {
		fn(key)
	}
}

func init() {
	cacheCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "show all data")
}
