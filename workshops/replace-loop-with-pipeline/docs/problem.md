# Replace Loop with Pipeline: Problem

## Goal

Refactor a loop-based function into a pipeline-style implementation without changing its behavior.

The starting code reads CSV-style office data, skips the header row, ignores blank lines, keeps only offices in India, and returns a simplified result with city and phone number.

## Learning Objective

This workshop teaches how to take a loop that does several different jobs and turn it into a sequence of smaller transformation steps.

In this exercise, the loop currently does all of the following:

- skips the header row
- ignores blank lines
- splits each CSV row into fields
- filters for India records
- maps records into output structs

The refactoring goal is to express those steps as a readable pipeline.

## Starting Point

The learner-facing code is in:

- `workshops/replace-loop-with-pipeline/example/exercise.go`

The function already works. This is intentional. Refactoring should preserve behavior while improving structure.

The pass/fail check is in:

- `workshops/replace-loop-with-pipeline/example/main_test.go`

## Your Task

Refactor `acquireIndianOfficeContacts` so that it:

- removes the main loop as the center of the logic
- replaces loop behavior with step-by-step collection transformations
- keeps the same output
- stays readable for a Go learner

## Go-Specific Note

Go does not have built-in method-chaining collection pipelines like some other languages. For this workshop, a pipeline is represented using small helper functions that transform slices in sequence.

That means the learner should focus on the refactoring principle, not on copying JavaScript or Java syntax directly.

## Suggested Refactoring Path

1. Identify the collection the loop iterates over.
2. Move header-skipping logic out of the loop.
3. Move blank-line removal into a filtering step.
4. Move string splitting into a mapping step.
5. Filter down to India records.
6. Map those records into `OfficeContact` values.
7. Replace the accumulator loop with the final pipeline result.

## Expected Output

When you run either the starting code or the finished solution, the output should be:

```text
Indian offices:
- Bangalore: +91 80 4064 9570
- Chennai: +91 44 660 44766
```

## Run Instructions

Run the starting exercise:

```bash
go run ./workshops/replace-loop-with-pipeline/example
```

Run the learner tests:

```bash
go test ./workshops/replace-loop-with-pipeline/example
```

The important constraint is that your refactor should keep the tests passing. The tests validate behavior, not whether you used the exact same pipeline structure as the reference solution.
