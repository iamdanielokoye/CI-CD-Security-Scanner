package main

import (
	"ci-cd-security-scanner/internal"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: scanner <config_file>")
		os.Exit(1)
	}
	configPath := os.Args[1]
	internal.RunScan(configPath)
}
