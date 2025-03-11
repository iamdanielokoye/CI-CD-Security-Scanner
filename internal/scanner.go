package internal

import (
	"fmt"
	"regexp"
	"strings"
)

// SecurityIssue represents a detected security issue.
type SecurityIssue struct {
	Severity string
	Message  string
	Line     int
}

// CheckInsecurePractices scans the CI/CD configuration for security issues.
func CheckInsecurePractices(configLines []string) []SecurityIssue {
	var issues []SecurityIssue
	secretPattern := regexp.MustCompile(`(?i)(password|token|apikey|secret)\s*=\s*['"][^'"]+['"]`)
	latestImagePattern := regexp.MustCompile(`image:\s*\S+:latest`)
	rootUserPattern := regexp.MustCompile(`user:\s*root`)

	for i, line := range configLines {
		trimmed := strings.TrimSpace(line)

		if secretPattern.MatchString(trimmed) {
			issues = append(issues, SecurityIssue{"CRITICAL", "Hardcoded secret detected", i + 1})
		}

		if latestImagePattern.MatchString(trimmed) {
			issues = append(issues, SecurityIssue{"WARNING", "Using 'latest' tag in Docker image", i + 1})
		}

		if rootUserPattern.MatchString(trimmed) {
			issues = append(issues, SecurityIssue{"HIGH", "Pipeline job runs as root", i + 1})
		}

		if strings.Contains(trimmed, "chmod 777") {
			issues = append(issues, SecurityIssue{"WARNING", "Overly permissive file permission (chmod 777)", i + 1})
		}
	}

	return issues
}

// PrintIssues displays security issues found in the configuration.
func PrintIssues(issues []SecurityIssue) {
	for _, issue := range issues {
		fmt.Printf("[%s] Line %d: %s\n", issue.Severity, issue.Line, issue.Message)
	}
}

func RunScan(configPath string) {
	config, err := ParseConfig(configPath)
	if err != nil {
		fmt.Println("Error parsing config:", err)
		return
	}

	EvaluateSecurity(config)
}
