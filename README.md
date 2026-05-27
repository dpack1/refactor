# Go Refactor Training, But By Grug

**Grug Brain** | [Big Brain](README.big-brain.md)

repo hold four small go workshop inspired by four refactoring moves from [Refactoring: Improving the Design of Existing Code](https://www.oreilly.com/library/view/refactoring-improving-the/9780134757681/cover.xhtml).

good book.
good for grug.
good for grug who want fewer code smell and fewer surprise bonk on head.

grug make small go workshop because large workshop often become side quest.
side quest bad when goal only learn few thing.

goal simple:

- teach one refactor
- keep workshop small
- let learner finish before cocoa juice get cold

each workshop have:

- `example/` where young grug make change
- `solution/` where older grug already bonk code into shape
- `docs/` where grug explain why bonk happen and where club land

repo for five minute workshop.
five minute good.
five minute enough to learn thing.
five minute not enough for complexity demon build full nest.

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
grug not at war with loop.
but sometimes small helper chain make shape of code easier see with tired cave eyes.
important part not "pipeline fancy".
important part "grug read code, grug know what code doing".

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

big conditional make grug squint and lean toward screen.
named function tell story sooner.
grug like when meaning come before forehead wrinkle.

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
then string collect rule.
then string become tiny chaos goblin in business disguise.
when this happen, grug give rule proper home.

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
grug drop plate.
query method let grug ask object for answer instead of juggling whole circus in one function.

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
many fence in old code there for reason.
