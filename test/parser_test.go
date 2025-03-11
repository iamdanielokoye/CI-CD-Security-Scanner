package test

import (
	"CI-CD-Security-Scanner/internal"
	"testing"
)

func TestParseConfig(t *testing.T) {
	_, err := internal.ParseConfig("../examples/secure-gitlab.yml")
	if err != nil {
		t.Errorf("Failed to parse valid CI/CD config: %v", err)
	}
}
