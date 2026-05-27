# Replace Loop with Pipeline: Solution Walkthrough

## Overview

The finished solution keeps the same behavior as the starting code, but it breaks the loop logic into a sequence of transformations.

The completed code lives in:

- `workshops/replace-loop-with-pipeline/solution/exercise.go`

The reference tests live in:

- `workshops/replace-loop-with-pipeline/solution/main_test.go`

## Step 1: Start with the input collection

The original function starts by splitting the CSV input into lines:

```go
lines := strings.Split(input, "\n")
```

This stays in place because it gives the rest of the function a clear collection to work from.

## Step 2: Remove the header outside the loop

The loop originally used a `firstLine` control variable to skip the CSV header. In the refactored version, the header is removed before later processing starts:

```go
lines[1:]
```

This is the Go equivalent of moving header-skipping behavior into the pipeline.

## Step 3: Filter out blank lines

Instead of checking for blank lines inside the loop with `continue`, the solution uses `Filter`:

```go
Filter(lines[1:], func(line string) bool {
	return strings.TrimSpace(line) != ""
})
```

This makes the intent more direct: only keep meaningful rows.

## Step 4: Map each line into fields

Each remaining CSV line is turned into a slice of fields:

```go
Map(..., func(line string) []string {
	return strings.Split(line, ",")
})
```

At this point, the data has moved from raw text rows to structured field slices.

## Step 5: Filter for India records

The next filter keeps only the records where the country column is `India`:

```go
Filter(..., func(fields []string) bool {
	return len(fields) >= 3 && strings.TrimSpace(fields[1]) == "India"
})
```

The `len(fields) >= 3` check keeps the pipeline safe if the input is malformed.

## Step 6: Map into the final output shape

The last transformation maps each matching record into an `OfficeContact`:

```go
Map(..., func(fields []string) OfficeContact {
	return OfficeContact{
		City:  strings.TrimSpace(fields[0]),
		Phone: strings.TrimSpace(fields[2]),
	}
})
```

This replaces the old accumulator pattern where the loop appended values one by one.

## Step 7: Return the pipeline result directly

Once all loop behavior has been moved into transformations, the function can return the final result directly:

```go
return Map(...)
```

That is the key outcome of the refactoring. The logic now reads as a flow of data:

1. drop the header
2. remove blank lines
3. split into fields
4. keep India rows
5. shape the result

## Why the Solution Uses Helper Functions

Go does not provide built-in `map` and `filter` operations on slices. To keep the exercise focused on the refactoring idea, the solution defines small generic helpers:

- `Filter[T any]`
- `Map[T any, U any]`

These helpers make the transformation pipeline explicit without bringing in extra libraries.

## Run Instructions

Run the completed solution:

```bash
go run ./workshops/replace-loop-with-pipeline/solution
```

Run the solution tests:

```bash
go test ./workshops/replace-loop-with-pipeline/solution
```
