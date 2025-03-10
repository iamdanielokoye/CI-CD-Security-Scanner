package internal

import (
	"fmt"
)

func EvaluateSecurity(config *CIConfig) {
	for job, details := range config.Jobs {
		fmt.Printf("Evaluating security for job: %s\n", job)
		fmt.Printf("Job details: %+v\n", details)
	}
	fmt.Println("Security evaluation complete.")
}
