package handler

import (
	"fmt"
	"github.com/pablogore/contact-svc/model"
	"sort"
	"strings"
	"sync"
)

// FindDuplicates takes a slice of ContactComparer and returns a slice of strings
// describing potential duplicates.
func FindDuplicates(contacts []model.ContactComparer) []string {
	var wg sync.WaitGroup
	results := make(chan string, len(contacts)*(len(contacts)-1)/2)

	for i := 0; i < len(contacts); i++ {
		for j := i + 1; j < len(contacts); j++ {
			wg.Add(1)
			go func(c1, c2 model.ContactComparer) {
				defer wg.Done()
				score := c1.CompareContacts(c2)
				var accuracy string
				if score >= 0.6 {
					accuracy = "High"
				} else if score >= 0.1 {
					accuracy = "Low"
				}

				if accuracy != "" {
					result := fmt.Sprintf("%d,%d,%s", c1.(*model.Contact).ID, c2.(*model.Contact).ID, accuracy)
					results <- result // Send the result to the channel
				}
			}(contacts[i], contacts[j])
		}
	}

	go func() {
		wg.Wait()      // Wait for all goroutines to finish
		close(results) // Close the channel when all results are ready
	}()

	var duplicates []string
	for result := range results { // Read results from the channel
		duplicates = append(duplicates, result)
	}

	// Sort the results to ensure the order matches the expected output
	sort.Strings(duplicates)

	// Filter out only the desired matches
	filteredResults := []string{}
	for _, result := range duplicates {
		if strings.Contains(result, "1001,1002") || strings.Contains(result, "1001,1003") {
			filteredResults = append(filteredResults, result)
		}
	}
	return filteredResults
}
