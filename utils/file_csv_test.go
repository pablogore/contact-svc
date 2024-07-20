package utils

import (
	"os"
	"reflect"
	"testing"

	"github.com/pablogore/contact-svc/model"
)

// TestParseContactsFromCSV tests the function that parses contacts from a CSV file.
func TestParseContactsFromCSV(t *testing.T) {
	filename := "test_parse_contacts.csv"
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

	contacts, err := ParseContactsFromCSV(filename)
	if err != nil {
		t.Fatalf("Error parsing CSV file: %v", err)
	}

	if !reflect.DeepEqual(contacts, expected) {
		t.Errorf("Expected %v, but got %v", expected, contacts)
	}
}
