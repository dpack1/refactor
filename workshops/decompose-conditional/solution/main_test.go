package main

import (
	"math"
	"testing"
	"time"
)

func TestCalculateCharge_UsesSummerRateInsideRange(t *testing.T) {
	date := time.Date(2026, time.July, 15, 0, 0, 0, 0, time.UTC)

	got := calculateCharge(date, 5, samplePlan)
	want := 4.25

	assertClose(t, got, want)
}

func TestCalculateCharge_UsesSummerRateOnBoundaries(t *testing.T) {
	start := calculateCharge(samplePlan.SummerStart, 3, samplePlan)
	end := calculateCharge(samplePlan.SummerEnd, 3, samplePlan)
	want := 2.55

	assertClose(t, start, want)
	assertClose(t, end, want)
}

func TestCalculateCharge_UsesRegularRateOutsideRange(t *testing.T) {
	date := time.Date(2026, time.November, 10, 0, 0, 0, 0, time.UTC)

	got := calculateCharge(date, 5, samplePlan)
	want := 18.25

	assertClose(t, got, want)
}

func assertClose(t *testing.T, got, want float64) {
	t.Helper()

	if math.Abs(got-want) > 0.00001 {
		t.Fatalf("unexpected charge: got %.2f want %.2f", got, want)
	}
}
