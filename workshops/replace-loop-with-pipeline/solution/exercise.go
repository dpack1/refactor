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

func acquireIndianOfficeContacts(input string) []OfficeContact {
	lines := strings.Split(input, "\n")
	if len(lines) <= 1 {
		return nil
	}

	return Map(
		Filter(
			Map(
				Filter(lines[1:], func(line string) bool {
					return strings.TrimSpace(line) != ""
				}),
				func(line string) []string {
					return strings.Split(line, ",")
				},
			),
			func(fields []string) bool {
				return len(fields) >= 3 && strings.TrimSpace(fields[1]) == "India"
			},
		),
		func(fields []string) OfficeContact {
			return OfficeContact{
				City:  strings.TrimSpace(fields[0]),
				Phone: strings.TrimSpace(fields[2]),
			}
		},
	)
}

func Filter[T any](items []T, keep func(T) bool) []T {
	result := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, U any](items []T, transform func(T) U) []U {
	result := make([]U, 0, len(items))
	for _, item := range items {
		result = append(result, transform(item))
	}
	return result
}
