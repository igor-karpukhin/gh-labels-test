package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func MatchWildcards(labels []string, testLabels []string, testType string) []string {
	matchedLabels := make(map[string]struct{})

	prefixMatch := regexp.MustCompile(fmt.Sprintf("^test/%s/(.+)$", testType))

	for _, label := range labels {
		if label == fmt.Sprintf("test/%s/*", testType) {
			for _, test := range testLabels {
				matchedLabels[test] = struct{}{}
			}
		} else {
			if prefixMatch.MatchString(label) {
				pattern := strings.Replace(prefixMatch.FindStringSubmatch(label)[1], "*", ".*", -1)
				for _, test := range testLabels {
					matched, _ := regexp.MatchString("^"+pattern+"$", test)
					if matched {
						matchedLabels[test] = struct{}{}
					}
				}
			}
		}
	}

	var result []string
	for key := range matchedLabels {
		result = append(result, key)
	}

	return result
}

func main() {
	envPRLabels := os.Getenv("PR_LABELS")
	envIntLabels := os.Getenv("INT_LABELS")
	envE2ELabels := os.Getenv("E2E_LABELS")

	var labels []string
	var intLabels []string
	var e2eLabels []string

	if err := json.Unmarshal([]byte(envPRLabels), &labels); err != nil {
		fmt.Printf("Error parsing labels: %v\n", err)
		return
	}
	if err := json.Unmarshal([]byte(envIntLabels), &intLabels); err != nil {
		fmt.Printf("Error parsing integration labels: %v\n", err)
		return
	}
	if err := json.Unmarshal([]byte(envE2ELabels), &e2eLabels); err != nil {
		fmt.Printf("Error parsing E2E labels: %v\n", err)
		return
	}

	matchedIntTests := MatchWildcards(labels, intLabels, "int")
	matchedE2ETests := MatchWildcards(labels, e2eLabels, "e2e")

	matchedIntTestsJSON, _ := json.Marshal(matchedIntTests)
	matchedE2ETestsJSON, _ := json.Marshal(matchedE2ETests)

	fmt.Printf("Matched Integration Tests: %s\n", matchedIntTestsJSON)
	fmt.Printf("Matched E2E Tests: %s\n", matchedE2ETestsJSON)
}
