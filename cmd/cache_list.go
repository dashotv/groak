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
			err := eachBucket(func(bucket string) error {
				log.Printf("-  %s", bucket)
				err := eachItem(bucket, func(key string) error {
					val, err := db.Get(bucket, key)
					if err != nil {
						return fmt.Errorf("%s failed to get value: %s", key, err)
					}
					log.Printf("   -  %s %s", key, val)
					return nil
				})
				if err != nil {
					log.Fatalf("error: %s\n", err)
				}
				return nil
			})
			if err != nil {
				log.Fatalf("error: %s\n", err)
			}
			return
		}

		if len(args) == 0 {
			_ = eachBucket(func(bucket string) error {
				log.Printf("-  %s", bucket)
				return nil
			})
		}

		bucket := args[0]
		err := eachItem(bucket, func(key string) error {
			val, err := db.Get(bucket, key)
			if err != nil {
				return fmt.Errorf("%s failed to get value: %s", key, err)
			}
			log.Printf("   -  %s %s", key, val)
			return nil
		})
		if err != nil {
			log.Fatalf("error: %s\n", err)
		}
	},
}

func eachBucket(fn func(string) error) error {
	buckets, err := db.ListBuckets()
	if err != nil {
		return fmt.Errorf("failed to list buckets: %s\n", err)
	}

	for _, bucket := range buckets {
		if bucket == "groak" {
			continue
		}
		err := fn(bucket)
		if err != nil {
			return err
		}
	}
	return nil
}

func eachItem(bucket string, fn func(string) error) error {
	keys, err := db.List(bucket)
	if err != nil {
		return fmt.Errorf("failed to list keys: %s\n", err)
	}

	for _, key := range keys {
		err := fn(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	cacheCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "show all data")
}
