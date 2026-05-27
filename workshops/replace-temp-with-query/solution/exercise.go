package main

type Order struct {
	quantity  int
	itemPrice float64
}

func NewOrder(quantity int, itemPrice float64) Order {
	return Order{quantity: quantity, itemPrice: itemPrice}
}

func (o Order) Price() float64 {
	return o.basePrice() * o.discountFactor()
}

func (o Order) basePrice() float64 {
	return float64(o.quantity) * o.itemPrice
}

func (o Order) discountFactor() float64 {
	discountFactor := 0.98

	if o.basePrice() > 1000 {
		discountFactor -= 0.03
	}

	return discountFactor
}

var sampleOrder = NewOrder(50, 25.00)
