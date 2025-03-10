package internal

import (
	"fmt"
)

func RunScan(configPath string) {
	config, err := ParseConfig(configPath)
	if err != nil {
		fmt.Println("Error parsing config:", err)
		return
	}

	EvaluateSecurity(config)
}
