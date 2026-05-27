# Replace Temp with Query: Problem

## Goal

Replace temporary variables in a method with small query methods on the same struct.

This workshop is intentionally small. The learner should be able to complete it in about five minutes.

## Edit This File

- `workshops/replace-temp-with-query/example/exercise.go`

## Your Task

Refactor `Order.Price()` so it no longer calculates `basePrice` and `discountFactor` as temporary variables inside the method.

Instead, move those calculations into query methods on `Order`.

Keep the behavior the same.

## Hint

The final `Price()` method should read more like:

```go
return o.basePrice() * o.discountFactor()
```

## Run

```bash
go test ./workshops/replace-temp-with-query/example
```
