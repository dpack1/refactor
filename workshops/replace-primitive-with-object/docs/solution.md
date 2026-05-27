# Replace Primitive with Object: Solution Walkthrough

## Overview

The solution introduces a `Priority` type so the `Order` type no longer owns raw string comparison logic.

The completed code lives in:

- `workshops/replace-primitive-with-object/solution/exercise.go`

## Step 1: Add a small value object

The solution adds a `Priority` type that wraps the string value:

```go
type Priority struct {
	value string
}
```

That gives the code a dedicated place for priority-specific behavior.

## Step 2: Move behavior onto `Priority`

Instead of letting `Order` compare raw strings directly, the logic moves into methods on `Priority`:

```go
func (p Priority) String() string
func (p Priority) NeedsAttention() bool
```

This is the key part of the refactoring. The primitive value now has a home for its own rules.

## Step 3: Update `Order` to use the new type

`Order` changes from:

```go
priority string
```

to:

```go
priority Priority
```

The constructor wraps the incoming string, and `Order` delegates back to the `Priority` object.

## Result

The public behavior stays the same, but the code is easier to extend later. If you later add validation, ranking, or display formatting, that logic has a clear place to live.

## Run

```bash
go test ./workshops/replace-primitive-with-object/solution
```
