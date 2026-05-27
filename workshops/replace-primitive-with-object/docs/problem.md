# Replace Primitive with Object: Problem

## Goal

Replace a primitive string priority with a small object that owns priority-specific behavior.

This workshop is intentionally small. The learner should be able to complete it in about five minutes.

## Edit This File

- `workshops/replace-primitive-with-object/example/exercise.go`

## Your Task

Refactor `Order` so it no longer stores priority as a plain string.

Create a small `Priority` type and move the priority logic onto it.

Keep these behaviors the same:

- `PriorityString()` still returns the original text value
- `NeedsAttention()` is true for `high` and `rush`
- `countOrdersNeedingAttention()` still returns the same count

## Hint

Keep the change small:

1. add a `Priority` type
2. store that type inside `Order`
3. delegate `PriorityString()` and `NeedsAttention()` to the new type

## Run

```bash
go test ./workshops/replace-primitive-with-object/example
```
