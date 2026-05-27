package main

import "time"

type PricingPlan struct {
	SummerStart          time.Time
	SummerEnd            time.Time
	SummerRate           float64
	RegularRate          float64
	RegularServiceCharge float64
}

var samplePlan = PricingPlan{
	SummerStart:          time.Date(2026, time.June, 1, 0, 0, 0, 0, time.UTC),
	SummerEnd:            time.Date(2026, time.August, 31, 0, 0, 0, 0, time.UTC),
	SummerRate:           0.85,
	RegularRate:          1.15,
	RegularServiceCharge: 12.50,
}

// calculateCharge is the starting point for the workshop.
// It works, but the condition and both branches need to be decomposed.
func calculateCharge(aDate time.Time, quantity int, plan PricingPlan) float64 {
	if (!aDate.Before(plan.SummerStart)) && (!aDate.After(plan.SummerEnd)) {
		return float64(quantity) * plan.SummerRate
	}

	return float64(quantity)*plan.RegularRate + plan.RegularServiceCharge
}
