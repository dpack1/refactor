# AGENTS.md

This file defines the working rules for building and maintaining Go training material in this directory. Update it over time as the workshop format becomes more opinionated.

## Purpose

This repository is used to create small Go workshops that include:

- a learner-facing problem or starting point
- a completed solution
- supporting documentation that explains how to move from the problem to the solution

The material may originate from examples written in another language, including Java, but the final training content in this repository must be Go-specific.

## Core Output Structure

Each workshop should aim to live in its own folder under `workshops/` and include:

- `example/` for the learner starting point
- `solution/` for the completed working answer
- `docs/` for walkthroughs, instructions, or teaching notes

Suggested file layout:

```text
workshops/
  workshop-name/
    example/
      main.go
    solution/
      main.go
    docs/
      problem.md
      solution.md
```

Adjust the structure when the exercise requires multiple files, packages, or tests, but keep the split between `example` and `solution` clear.

## Authoring Rules

- Keep exercises small and teach one primary idea at a time.
- Prefer simple, readable Go over clever abstractions.
- Use standard library solutions unless a dependency is necessary for the lesson.
- Make the example intentionally incomplete, incorrect, or in need of refactoring so the learner has a clear task.
- Make the solution fully working and easy to compare with the example.
- Keep naming explicit and beginner-friendly.
- Avoid unnecessary framework setup unless the exercise is specifically about tooling or architecture.

## Translation Rules

When adapting source material from Java or another language:

- translate the intent, not the syntax
- rewrite the example so it feels native to Go
- remove patterns that only make sense in Java
- use idiomatic Go error handling, structs, interfaces, and package layout where appropriate
- keep the learning goal equivalent to the original source

## Documentation Rules

For each workshop, provide two docs when possible:

- a problem doc that explains the task, constraints, and expected outcome
- a solution doc that explains the implementation path step by step

Documentation should:

- assume the learner may be new to Go
- explain why the refactor or fix matters
- call out important Go concepts directly
- include run instructions when relevant
- avoid large theory sections that are not needed for the exercise

## Quality Bar

Before considering a workshop complete:

- the solution should run successfully
- the example should be understandable and usable as an exercise
- the docs should clearly connect the problem to the solution
- file names and folder names should be predictable
- the exercise should be completable without hidden setup

## Maintenance

As this repo evolves, extend this file with:

- naming conventions
- preferred workshop templates
- testing expectations
- review checklist items
- style decisions for docs and code

## Naming Conventions

- Use lowercase, hyphenated folder names for workshops, such as `replace-loop-with-pipeline`.
- Keep each workshop runnable from its own `example/` and `solution/` directory.
- Prefer `main.go` for small workshop entry points unless the lesson specifically needs multiple packages.

## 5-Minute Workshop Guidelines

- Design each workshop so a learner can complete it in about five minutes.
- Keep the exercise centered on one refactoring move, not on setup or domain complexity.
- Prefer one small file to edit, plus tests.
- Keep the starting code around 8 to 15 lines of real logic when possible.
- Use 2 or 3 focused tests instead of a large test suite.
- Avoid parsing, file I/O, dates, and extra domain rules unless they are essential to the refactoring itself.
- Keep docs short: goal, file to edit, test command, and one or two hints.
