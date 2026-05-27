package main

type Order struct {
	quantity  int
	itemPrice float64
}

func NewOrder(quantity int, itemPrice float64) Order {
	return Order{quantity: quantity, itemPrice: itemPrice}
}

// Price is the starting point for the workshop.
// It works, but the temporary variables should be replaced with query methods.
func (o Order) Price() float64 {
	basePrice := float64(o.quantity) * o.itemPrice
	discountFactor := 0.98

	if basePrice > 1000 {
		discountFactor -= 0.03
	}

	return basePrice * discountFactor
}

var sampleOrder = NewOrder(50, 25.00)
