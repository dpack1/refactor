# Decompose Conditional: Solution Walkthrough

## Overview

The completed solution keeps the same pricing behavior, but it makes the decision easier to read by extracting the condition and both branches into named functions.

The completed code lives in:

- `workshops/decompose-conditional/solution/exercise.go`

The reference tests live in:

- `workshops/decompose-conditional/solution/main_test.go`

## Step 1: Extract the condition

The original function hides the reason for the branch inside a date comparison:

```go
!aDate.Before(plan.SummerStart) && !aDate.After(plan.SummerEnd)
```

That logic is moved into `isSummer`, which gives the condition a clear name:

```go
func isSummer(aDate time.Time, plan PricingPlan) bool {
	return !aDate.Before(plan.SummerStart) && !aDate.After(plan.SummerEnd)
}
```

## Step 2: Extract the summer branch

The `if` branch becomes a focused helper:

```go
func summerCharge(quantity int, plan PricingPlan) float64 {
	return float64(quantity) * plan.SummerRate
}
```

This makes the main function read in terms of business meaning instead of arithmetic details.

## Step 3: Extract the regular branch

The `else` branch also becomes a focused helper:

```go
func regularCharge(quantity int, plan PricingPlan) float64 {
	return float64(quantity)*plan.RegularRate + plan.RegularServiceCharge
}
```

Now both pricing paths have explicit names.

## Step 4: Simplify the original function

Once the condition and both branches have names, the original function becomes:

```go
func calculateCharge(aDate time.Time, quantity int, plan PricingPlan) float64 {
	if isSummer(aDate, plan) {
		return summerCharge(quantity, plan)
	}

	return regularCharge(quantity, plan)
}
```

That is the core of the refactoring. The function now shows:

1. what decision is being made
2. what happens in the summer case
3. what happens otherwise

## Why This Refactoring Helps

The original version tells you the mechanics of the rule, but not the intent. The decomposed version exposes the business idea directly:

- `isSummer`
- `summerCharge`
- `regularCharge`

That makes the code easier to scan, easier to change, and easier to discuss with another developer.

## Run Instructions

Run the completed solution:

```bash
go run ./workshops/decompose-conditional/solution
```

Run the solution tests:

```bash
go test ./workshops/decompose-conditional/solution
```
