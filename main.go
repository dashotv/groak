/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/dotenv-org/godotenvvault"

	"github.com/dashotv/groak/cmd"
)

func main() {
	// Load .env file and populate environment variables
	if err := godotenvvault.Load(); err != nil {
		fmt.Printf("failed to load env: %s\n", err)
		os.Exit(1)
	}
	cmd.Execute()
}
