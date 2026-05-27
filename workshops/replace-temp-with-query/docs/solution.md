# Replace Temp with Query: Solution Walkthrough

## Overview

The solution moves temporary calculations out of `Price()` and into small query methods on `Order`.

The completed code lives in:

- `workshops/replace-temp-with-query/solution/exercise.go`

## Step 1: Extract the base price calculation

The original `Price()` method calculated:

```go
basePrice := float64(o.quantity) * o.itemPrice
```

This becomes:

```go
func (o Order) basePrice() float64 {
	return float64(o.quantity) * o.itemPrice
}
```

## Step 2: Extract the discount factor calculation

The second temp also moves into its own query method:

```go
func (o Order) discountFactor() float64 {
	discountFactor := 0.98
	if o.basePrice() > 1000 {
		discountFactor -= 0.03
	}
	return discountFactor
}
```

## Step 3: Simplify `Price()`

Once both calculations have names, `Price()` becomes:

```go
func (o Order) Price() float64 {
	return o.basePrice() * o.discountFactor()
}
```

That is the core of the refactoring. The temporary values are gone, but the meaning of the code is clearer.

## Run

```bash
go test ./workshops/replace-temp-with-query/solution
```
