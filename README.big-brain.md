# Go Refactor Training

[Grug Brain](README.md) | **Big Brain**

This repository highlights four small Go workshops inspired by four refactoring moves from [Refactoring: Improving the Design of Existing Code](https://www.oreilly.com/library/view/refactoring-improving-the/9780134757681/cover.xhtml).

Small Go workshops for practicing one refactoring move at a time in idiomatic Go.

Each workshop includes:

- `example/` for the learner starting point
- `solution/` for the completed answer
- `docs/` for the problem statement and walkthrough

The repository is organized for short exercises that should be completable in about five minutes and focus on Go structs, methods, helpers, and slice-oriented code.

## Workshops

### [Replace Loop with Pipeline](workshops/replace-loop-with-pipeline/)

<table>
  <tr>
    <th>Before</th>
    <th>After</th>
  </tr>
  <tr>
    <td>
      <pre lang="go"><code>names := make([]string, 0, len(users))
for _, user := range users {
    if user.Active {
        names = append(names, user.Name)
    }
}</code></pre>
    </td>
    <td>
      <pre lang="go"><code>names := Map(
    Filter(users, isActive),
    nameOf,
)</code></pre>
    </td>
  </tr>
</table>

Refactor a loop-based transformation into a small slice pipeline built from Go helper functions.

- Problem: [docs/problem.md](workshops/replace-loop-with-pipeline/docs/problem.md)
- Solution walkthrough: [docs/solution.md](workshops/replace-loop-with-pipeline/docs/solution.md)
- Starting code: [example/](workshops/replace-loop-with-pipeline/example/)
- Reference solution: [solution/](workshops/replace-loop-with-pipeline/solution/)

### [Decompose Conditional](workshops/decompose-conditional/)

<table>
  <tr>
    <th>Before</th>
    <th>After</th>
  </tr>
  <tr>
    <td>
      <pre lang="go"><code>if customer.isPremium && order.total() &gt; 100 {
    return order.total() * 0.9
}
return order.total()</code></pre>
    </td>
    <td>
      <pre lang="go"><code>if qualifiesForDiscount(customer, order) {
    return discountedTotal(order)
}
return fullTotal(order)</code></pre>
    </td>
  </tr>
</table>

Break a dense conditional into named Go functions so the rule and both branches read clearly.

- Problem: [docs/problem.md](workshops/decompose-conditional/docs/problem.md)
- Solution walkthrough: [docs/solution.md](workshops/decompose-conditional/docs/solution.md)
- Starting code: [example/](workshops/decompose-conditional/example/)
- Reference solution: [solution/](workshops/decompose-conditional/solution/)

### [Replace Primitive with Object](workshops/replace-primitive-with-object/)

<table>
  <tr>
    <th>Before</th>
    <th>After</th>
  </tr>
  <tr>
    <td>
      <pre lang="go"><code>type Ticket struct {
    severity string
}

func (t Ticket) NeedsPaging() bool {
    return t.severity == "critical"
}</code></pre>
    </td>
    <td>
      <pre lang="go"><code>type Severity struct {
    value string
}

func (t Ticket) NeedsPaging() bool {
    return t.severity.NeedsPaging()
}</code></pre>
    </td>
  </tr>
</table>

Replace a raw string field with a small Go type and move behavior onto that type.

- Problem: [docs/problem.md](workshops/replace-primitive-with-object/docs/problem.md)
- Solution walkthrough: [docs/solution.md](workshops/replace-primitive-with-object/docs/solution.md)
- Starting code: [example/](workshops/replace-primitive-with-object/example/)
- Reference solution: [solution/](workshops/replace-primitive-with-object/solution/)

### [Replace Temp with Query](workshops/replace-temp-with-query/)

<table>
  <tr>
    <th>Before</th>
    <th>After</th>
  </tr>
  <tr>
    <td>
      <pre lang="go"><code>subtotal := job.hours * job.hourlyRate
fee := subtotal * 0.12

if subtotal &gt; 500 {
    fee *= 0.5
}

return subtotal + fee</code></pre>
    </td>
    <td>
      <pre lang="go"><code>return job.subtotal() + job.fee()</code></pre>
    </td>
  </tr>
</table>

Replace temporary variables with small query methods on a Go struct.

- Problem: [docs/problem.md](workshops/replace-temp-with-query/docs/problem.md)
- Solution walkthrough: [docs/solution.md](workshops/replace-temp-with-query/docs/solution.md)
- Starting code: [example/](workshops/replace-temp-with-query/example/)
- Reference solution: [solution/](workshops/replace-temp-with-query/solution/)

## Quick Start

Run a workshop from its `example/` or `solution/` directory path:

```bash
go run ./workshops/replace-loop-with-pipeline/example
go test ./workshops/replace-loop-with-pipeline/example

go run ./workshops/replace-loop-with-pipeline/solution
go test ./workshops/replace-loop-with-pipeline/solution
```

Swap the workshop path to run any other exercise:

- `./workshops/decompose-conditional/...`
- `./workshops/replace-primitive-with-object/...`
- `./workshops/replace-temp-with-query/...`

## Repository Layout

```text
workshops/
  workshop-name/
    example/
    solution/
    docs/
```

Repository-specific authoring guidance lives in [AGENTS.md](AGENTS.md).
