package main

import (
	"math"
	"testing"
)

func TestOrderPrice_BelowDiscountThreshold(t *testing.T) {
	order := NewOrder(10, 50.00)

	got := order.Price()
	want := 490.00

	assertClose(t, got, want)
}

func TestOrderPrice_AboveDiscountThreshold(t *testing.T) {
	order := NewOrder(50, 25.00)

	got := order.Price()
	want := 1187.50

	assertClose(t, got, want)
}

func TestOrderPrice_AtThresholdUsesRegularDiscountFactor(t *testing.T) {
	order := NewOrder(20, 50.00)

	got := order.Price()
	want := 980.00

	assertClose(t, got, want)
}

func assertClose(t *testing.T, got, want float64) {
	t.Helper()

	if math.Abs(got-want) > 0.00001 {
		t.Fatalf("unexpected price: got %.2f want %.2f", got, want)
	}
}
