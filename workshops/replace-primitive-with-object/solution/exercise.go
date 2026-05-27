package main

type Priority struct {
	value string
}

func NewPriority(value string) Priority {
	return Priority{value: value}
}

func (p Priority) String() string {
	return p.value
}

func (p Priority) NeedsAttention() bool {
	return p.value == "high" || p.value == "rush"
}

type Order struct {
	priority Priority
}

func NewOrder(priority string) Order {
	return Order{priority: NewPriority(priority)}
}

func (o Order) PriorityString() string {
	return o.priority.String()
}

func (o Order) NeedsAttention() bool {
	return o.priority.NeedsAttention()
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
