# Go Refactor Training

This directory is for building small Go training workshops from source examples, including examples that originally came from Java.

## Navigation

1. [Overview](#go-refactor-training)
2. [Repository Layout](#repository-layout)
3. [Workshops](#workshops)
4. [Intended Workflow](#intended-workflow)
5. [Running a Workshop](#running-a-workshop)
6. [Notes](#notes)

The goal is to keep each workshop simple and practical:

- each workshop has an `example/` directory for the learner-facing starting point
- each workshop has a `solution/` directory for the completed working version
- each workshop has a `docs/` directory for the exercise and solution walkthroughs
- each workshop should be small enough to teach in about five minutes

## Repository Layout

```text
.
├── AGENTS.md
├── README.md
├── go.mod
└── workshops/
    ├── decompose-conditional/
    │   ├── docs/
    │   ├── example/
    │   └── solution/
    ├── replace-primitive-with-object/
    │   ├── docs/
    │   ├── example/
    │   └── solution/
    ├── replace-temp-with-query/
    │   ├── docs/
    │   ├── example/
    │   └── solution/
    └── replace-loop-with-pipeline/
        ├── docs/
        ├── example/
        └── solution/
```

This layout is intentionally simple so each workshop is self-contained and easy to browse.

## Workshops

### Replace Loop with Pipeline

- Path: `workshops/replace-loop-with-pipeline/`
- Focus: refactor a loop-based CSV transformation into a pipeline-style flow in Go
- Includes learner code, solution code, and tests for both sides

### Decompose Conditional

- Path: `workshops/decompose-conditional/`
- Focus: refactor a dense conditional into named helpers for the condition and both branches
- Includes learner code, solution code, and tests for both sides

### Replace Primitive with Object

- Path: `workshops/replace-primitive-with-object/`
- Focus: replace a string priority with a small `Priority` object and move behavior onto it
- Includes learner code, solution code, and tests for both sides

### Replace Temp with Query

- Path: `workshops/replace-temp-with-query/`
- Focus: replace temporary variables in a pricing method with query methods on the struct
- Includes learner code, solution code, and tests for both sides

## Intended Workflow

1. Start from a source example or teaching prompt.
2. Convert the exercise so it is idiomatic and teachable in Go.
3. Create the workshop folder under `workshops/`.
4. Create the learner-facing version in `example/`.
5. Create the finished version in `solution/`.
6. Write supporting docs that explain the task and the solution process.

## Running a Workshop

Run the starting version:

```bash
go run ./workshops/replace-loop-with-pipeline/example
```

Run the learner tests:

```bash
go test ./workshops/replace-loop-with-pipeline/example
```

Run the completed solution:

```bash
go run ./workshops/replace-loop-with-pipeline/solution
```

Run the reference solution tests:

```bash
go test ./workshops/replace-loop-with-pipeline/solution
```

Run the second workshop:

```bash
go run ./workshops/decompose-conditional/example
go test ./workshops/decompose-conditional/example
go run ./workshops/decompose-conditional/solution
go test ./workshops/decompose-conditional/solution
```

Run the third workshop:

```bash
go run ./workshops/replace-primitive-with-object/example
go test ./workshops/replace-primitive-with-object/example
go run ./workshops/replace-primitive-with-object/solution
go test ./workshops/replace-primitive-with-object/solution
```

Run the fourth workshop:

```bash
go run ./workshops/replace-temp-with-query/example
go test ./workshops/replace-temp-with-query/example
go run ./workshops/replace-temp-with-query/solution
go test ./workshops/replace-temp-with-query/solution
```

## Notes

- Repository guidance lives in `AGENTS.md`.
- Each workshop should keep the split between problem and solution obvious.
