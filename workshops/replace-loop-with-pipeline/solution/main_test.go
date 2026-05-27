package main

import (
	"reflect"
	"testing"
)

func TestAcquireIndianOfficeContacts_ReturnsOnlyIndianOffices(t *testing.T) {
	got := acquireIndianOfficeContacts(officeData)

	want := []OfficeContact{
		{City: "Bangalore", Phone: "+91 80 4064 9570"},
		{City: "Chennai", Phone: "+91 44 660 44766"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected contacts: got %#v want %#v", got, want)
	}
}

func TestAcquireIndianOfficeContacts_IgnoresBlankAndMalformedRows(t *testing.T) {
	input := `office,country,telephone

Mumbai,India,+91 22 1234 5678
Broken Row
Berlin,Germany,+49 30 123456
Pune,India,+91 20 7654 3210
`

	got := acquireIndianOfficeContacts(input)

	want := []OfficeContact{
		{City: "Mumbai", Phone: "+91 22 1234 5678"},
		{City: "Pune", Phone: "+91 20 7654 3210"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected contacts: got %#v want %#v", got, want)
	}
}

func TestAcquireIndianOfficeContacts_EmptyInputReturnsNoContacts(t *testing.T) {
	got := acquireIndianOfficeContacts("office,country,telephone\n")

	if len(got) != 0 {
		t.Fatalf("expected no contacts, got %#v", got)
	}
}
