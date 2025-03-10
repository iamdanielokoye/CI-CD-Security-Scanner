package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type CIConfig struct {
	Jobs map[string]interface{} `yaml:"jobs"`
}

func ParseConfig(filePath string) (*CIConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config CIConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	fmt.Println("Parsed CI/CD configuration successfully.")
	return &config, nil
}
