package main

import (
	"fmt"
	"github.com/pablogore/contact-svc/handler"
	"github.com/pablogore/contact-svc/utils"
)

func main() {
	// Read contacts from the CSV file
	contacts, err := utils.ParseContactsFromCSV("contacts.csv")
	if err != nil {
		fmt.Println("Error parsing CSV:", err)
		return
	}

	// Find duplicates
	duplicates := handler.FindDuplicates(contacts)

	// Print results in the expected format
	fmt.Println("ContactID Source,ContactID Match,Accuracy")
	for _, result := range duplicates {
		fmt.Println(result)
	}
}
