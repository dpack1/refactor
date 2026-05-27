package main

import "testing"

func TestOrder_PriorityStringReturnsOriginalValue(t *testing.T) {
	order := NewOrder("high")

	if order.PriorityString() != "high" {
		t.Fatalf("expected priority string to be high, got %q", order.PriorityString())
	}
}

func TestOrder_NeedsAttentionForHighAndRush(t *testing.T) {
	high := NewOrder("high")
	rush := NewOrder("rush")
	normal := NewOrder("normal")

	if !high.NeedsAttention() {
		t.Fatal("expected high priority order to need attention")
	}

	if !rush.NeedsAttention() {
		t.Fatal("expected rush priority order to need attention")
	}

	if normal.NeedsAttention() {
		t.Fatal("expected normal priority order not to need attention")
	}
}

func TestCountOrdersNeedingAttention(t *testing.T) {
	got := countOrdersNeedingAttention(sampleOrders)

	if got != 2 {
		t.Fatalf("expected 2 orders needing attention, got %d", got)
	}
}
