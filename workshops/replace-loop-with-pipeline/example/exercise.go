package main

import "strings"

type OfficeContact struct {
	City  string
	Phone string
}

const officeData = `office,country,telephone
Chicago,USA,+1 312 373 1000
Beijing,China,+86 4008 900 505
Bangalore,India,+91 80 4064 9570
Porto Alegre,Brazil,+55 51 3079 3550
Chennai,India,+91 44 660 44766
`

// acquireIndianOfficeContacts is the starting point for the workshop.
// It works, but uses a loop and control flow that the learner should refactor.
func acquireIndianOfficeContacts(input string) []OfficeContact {
	lines := strings.Split(input, "\n")
	firstLine := true
	result := []OfficeContact{}

	for _, line := range lines {
		if firstLine {
			firstLine = false
			continue
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if len(fields) < 3 {
			continue
		}

		if strings.TrimSpace(fields[1]) == "India" {
			result = append(result, OfficeContact{
				City:  strings.TrimSpace(fields[0]),
				Phone: strings.TrimSpace(fields[2]),
			})
		}
	}

	return result
}
