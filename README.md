# Go Refactor Training, But By Grug

**Grug Brain** | [Big Brain](README.big-brain.md)

grug make small go workshop

goal simple:

- teach one refactor
- keep workshop small
- let learner finish before cocoa juice get cold

each workshop have:

- `example/` where young grug make change
- `solution/` where older grug already bonk complexity
- `docs/` where grug explain why bonk happen

repo for five minute workshop.
five minute good.
five minute leave less time for complexity demon to build nest.

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

loop fine.
but sometimes helper function make intent easier for grug brain.
important part is not "pipeline fancy".
important part is "can grug read thing without headache".

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

big conditional make grug squint.
named function help grug know what code mean before grug know every detail.
this good trade.

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

string start innocent.
string later become tiny chaos goblin.
when behavior cling to primitive, grug sometimes give primitive proper home.

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

temp variable sometimes helpful.
too many temp variable become plate spinning act.
query method let grug ask object for answer instead of juggling math in one place.

- Problem: [docs/problem.md](workshops/replace-temp-with-query/docs/problem.md)
- Solution walkthrough: [docs/solution.md](workshops/replace-temp-with-query/docs/solution.md)
- Starting code: [example/](workshops/replace-temp-with-query/example/)
- Reference solution: [solution/](workshops/replace-temp-with-query/solution/)

## Quick Start

run workshop like this:

```bash
go run ./workshops/replace-loop-with-pipeline/example
go test ./workshops/replace-loop-with-pipeline/example

go run ./workshops/replace-loop-with-pipeline/solution
go test ./workshops/replace-loop-with-pipeline/solution
```

change path when grug want other workshop:

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

more repo rule live in [AGENTS.md](AGENTS.md).
grug recommend read rule before wild refactor club swinging.
