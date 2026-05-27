package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2026, time.July, 15, 0, 0, 0, 0, time.UTC)
	charge := calculateCharge(date, 5, samplePlan)

	fmt.Printf("Charge for %s: $%.2f\n", date.Format("2006-01-02"), charge)
}
