package model

import (
	"strings"
)

// Contact represents a contact with basic details.
type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	ZipCode   string
	Address   string
}

// ContactComparer is an interface that defines the methods for comparing contacts.
type ContactComparer interface {
	CompareContacts(ContactComparer) float64
}

// CompareContacts compares the current Contact with another Contact
// and returns a similarity score based on weighted fields.
func (c1 *Contact) CompareContacts(c2 ContactComparer) float64 {
	c2Contact := c2.(*Contact) // Type assertion

	weights := map[string]float64{
		"FirstName": 0.2,
		"LastName":  0.2,
		"Email":     0.4,
		"ZipCode":   0.1,
		"Address":   0.1,
	}

	score := 0.0
	if strings.EqualFold(c1.FirstName, c2Contact.FirstName) {
		score += weights["FirstName"]
	}
	if strings.EqualFold(c1.LastName, c2Contact.LastName) {
		score += weights["LastName"]
	}
	if c1.Email == c2Contact.Email {
		score += weights["Email"]
	}
	if c1.ZipCode == c2Contact.ZipCode {
		score += weights["ZipCode"]
	}
	if c1.Address == c2Contact.Address {
		score += weights["Address"]
	}

	return score
}
