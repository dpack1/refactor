package main

type Order struct {
	priority string
}

func NewOrder(priority string) Order {
	return Order{priority: priority}
}

func (o Order) PriorityString() string {
	return o.priority
}

func (o Order) NeedsAttention() bool {
	return o.priority == "high" || o.priority == "rush"
}

func countOrdersNeedingAttention(orders []Order) int {
	count := 0
	for _, order := range orders {
		if order.NeedsAttention() {
			count++
		}
	}
	return count
}

var sampleOrders = []Order{
	NewOrder("low"),
	NewOrder("normal"),
	NewOrder("high"),
	NewOrder("rush"),
}
