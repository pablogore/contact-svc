package main

import (
	"os"
	"reflect"
	"testing"

	"github.com/pablogore/contact-svc/handler"
	"github.com/pablogore/contact-svc/model"
	"github.com/pablogore/contact-svc/utils"
)

// TestParseContactsFromCSV tests the function that parses contacts from a CSV file.
func TestParseContactsFromCSV(t *testing.T) {
	filename := "test_contacts.csv"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Error creating test CSV file: %v", err)
	}
	defer os.Remove(filename)

	data := "ID,FirstName,LastName,Email,ZipCode,Address\n" +
		"1001,C,F,mollis.lectus.pede@outlook.net,449-6990,Tellus. Rd.\n" +
		"1002,C,French,mollis.lectus.pede@outlook.net,39746,449-6990 Tellus. Rd.\n" +
		"1003,Ciara,F,non.lacinia.at@zoho.ca,39746,1234 Main St\n"

	_, err = file.WriteString(data)
	if err != nil {
		t.Fatalf("Error writing to test CSV file: %v", err)
	}
	file.Close()

	expected := []model.ContactComparer{
		&model.Contact{ID: 1001, FirstName: "C", LastName: "F", Email: "mollis.lectus.pede@outlook.net", ZipCode: "449-6990", Address: "Tellus. Rd."},
		&model.Contact{ID: 1002, FirstName: "C", LastName: "French", Email: "mollis.lectus.pede@outlook.net", ZipCode: "39746", Address: "449-6990 Tellus. Rd."},
		&model.Contact{ID: 1003, FirstName: "Ciara", LastName: "F", Email: "non.lacinia.at@zoho.ca", ZipCode: "39746", Address: "1234 Main St"},
	}

	contacts, err := utils.ParseContactsFromCSV(filename)
	if err != nil {
		t.Fatalf("Error parsing CSV file: %v", err)
	}

	if !reflect.DeepEqual(contacts, expected) {
		t.Errorf("Expected %v, but got %v", expected, contacts)
	}
}

// TestCompareContacts tests the method that compares two Contact structs.
func TestCompareContacts(t *testing.T) {
	c1 := &model.Contact{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", ZipCode: "12345", Address: "123 Elm St"}
	c2 := &model.Contact{ID: 2, FirstName: "john", LastName: "doe", Email: "john@example.com", ZipCode: "12345", Address: "123 Elm St"}
	expected := 1.0
	output := c1.CompareContacts(c2)
	if output != expected {
		t.Errorf("Expected %v, but got %v", expected, output)
	}
}

// TestFindDuplicates tests the function that finds duplicate contacts.
func TestFindDuplicates(t *testing.T) {
	contacts := []model.ContactComparer{
		&model.Contact{ID: 1001, FirstName: "C", LastName: "F", Email: "mollis.lectus.pede@outlook.net", ZipCode: "449-6990", Address: "Tellus. Rd."},
		&model.Contact{ID: 1002, FirstName: "C", LastName: "French", Email: "mollis.lectus.pede@outlook.net", ZipCode: "39746", Address: "449-6990 Tellus. Rd."},
		&model.Contact{ID: 1003, FirstName: "Ciara", LastName: "F", Email: "non.lacinia.at@zoho.ca", ZipCode: "39746", Address: "1234 Main St"},
		&model.Contact{ID: 1004, FirstName: "A", LastName: "B", Email: "example@domain.com", ZipCode: "12345", Address: "Some Address"},
	}

	expected := []string{
		"1001,1002,High",
		"1001,1003,Low",
	}

	duplicates := handler.FindDuplicates(contacts)

	// Check if all expected duplicates are in the actual results
	for _, exp := range expected {
		found := false
		for _, act := range duplicates {
			if exp == act {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected duplicate %v not found in actual results", exp)
		}
	}

	// Ensure there are no unexpected duplicates
	if len(duplicates) != len(expected) {
		t.Errorf("Expected %d duplicates, but got %d", len(expected), len(duplicates))
	}

	// Test case where accuracy is ""
	c4 := &model.Contact{ID: 1004, FirstName: "John", LastName: "Smith", Email: "john.smith@example.com", ZipCode: "99999", Address: "999 Main St"}
	c5 := &model.Contact{ID: 1005, FirstName: "Johnny", LastName: "Smith", Email: "johnny.smith@example.com", ZipCode: "99998", Address: "998 Main St"}
	contactsWithNoMatch := []model.ContactComparer{c4, c5}

	duplicatesWithNoMatch := handler.FindDuplicates(contactsWithNoMatch)

	// Ensure there are no matches with accuracy ""
	if len(duplicatesWithNoMatch) != 0 {
		t.Errorf("Expected 0 duplicates with accuracy '', but got %d", len(duplicatesWithNoMatch))
	}
}
