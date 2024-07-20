package utils

import (
	"encoding/csv"
	"github.com/pablogore/contact-svc/model"
	"os"
	"strconv"
)

// ParseContactsFromCSV reads a CSV file and parses its content into a slice of ContactComparer.
//
// The CSV file should have the following header:
// ID,FirstName,LastName,Email,ZipCode,Address
func ParseContactsFromCSV(filename string) ([]model.ContactComparer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var contacts []model.ContactComparer
	for _, record := range records[1:] { // Skip the header
		id, _ := strconv.Atoi(record[0])
		contacts = append(contacts, &model.Contact{
			ID:        id,
			FirstName: record[1],
			LastName:  record[2],
			Email:     record[3],
			ZipCode:   record[4],
			Address:   record[5],
		})
	}
	return contacts, nil
}
