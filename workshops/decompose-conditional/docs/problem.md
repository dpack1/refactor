# Decompose Conditional: Problem

## Goal

Refactor a function with a dense conditional into smaller, intention-revealing functions without changing its behavior.

The starting code calculates a charge using one rate during summer and another rate during the rest of the year.

## Learning Objective

This workshop teaches how to make conditional logic easier to read by extracting:

- the condition being checked
- the logic in the `if` branch
- the logic in the `else` branch

The original function works, but it mixes the rule, the branching decision, and the pricing calculations in one place.

## Starting Point

The learner-facing code is in:

- `workshops/decompose-conditional/example/exercise.go`

The pass/fail check is in:

- `workshops/decompose-conditional/example/main_test.go`

## Your Task

Refactor `calculateCharge` so that it:

- removes the dense inline condition from the main function
- extracts the summer check into a named helper
- extracts the summer calculation into a named helper
- extracts the regular calculation into a named helper
- keeps the same behavior

## Go-Specific Note

Go does not have a ternary operator, so the final code should keep a normal `if` statement. The point of the refactoring is still the same: make the branch easier to understand by naming the condition and the two outcomes.

## Suggested Refactoring Path

1. Extract the seasonal check into a helper such as `isSummer`.
2. Extract the summer calculation into a helper such as `summerCharge`.
3. Extract the non-summer calculation into a helper such as `regularCharge`.
4. Update `calculateCharge` so it reads as a simple high-level decision.

## Expected Behavior

The logic should produce these results:

- a date inside the summer window uses the summer rate
- the start and end dates of summer are included in the summer window
- a date outside the summer window uses the regular rate plus service charge

## Run Instructions

Run the starting exercise:

```bash
go run ./workshops/decompose-conditional/example
```

Run the learner tests:

```bash
go test ./workshops/decompose-conditional/example
```

The tests validate behavior. They do not enforce exact function names, but the refactoring should still clearly decompose the conditional.
