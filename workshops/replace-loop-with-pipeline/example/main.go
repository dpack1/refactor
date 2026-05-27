package main

import "fmt"

func main() {
	contacts := acquireIndianOfficeContacts(officeData)

	fmt.Println("Indian offices:")
	for _, contact := range contacts {
		fmt.Printf("- %s: %s\n", contact.City, contact.Phone)
	}
}
