# Challenge 5

## Challenge 5a

### Instructions

A flight booking platform displays search results as route strings. Each route shows the origin airport code, the destination, and the price.

Write a `RenderRoute` function that formats a flight into a display string for customers.

### Rules

- Trim origin and destination before use
- If origin or destination is empty after trim, return `""`
- Format price with exactly two decimal places

Return a string in this format: `HEL -> LHR (129.90 EUR)`

### Allowed packages

The entire Go standard library.

### Expected function

```go
package challenge

func RenderRoute(origin string, destination string, priceEUR float64) string {
    // solution
}
```

### Examples

```go
RenderRoute("HEL", "LHR", 129.90)
// -> "HEL -> LHR (129.90 EUR)"

RenderRoute(" HEL ", " LHR ", 0.99)
// -> "HEL -> LHR (0.99 EUR)"

RenderRoute(" ", "LHR", 12.00)
// -> ""

RenderRoute("HEL", "LHR", 0)
// -> "HEL -> LHR (0.00 EUR)"
```

---

## Challenge 5b

### Instructions

Gate staff reads boarding lines on an internal screen as passengers check in. A downstream display system parses the `|` separators to lay out columns.

Write a `RenderBoardingLines` function that formats boarding pass records into display lines for gate staff.

Each output line has the format: `<Name> | Seat <Seat> | Gate <Gate>`

### Rules

- Trim spaces in each field
- Skip rows where any field is empty after trim
- Preserve order of valid rows
- Use exactly one space on each side of `|`
- Pad each column to the width of the longest value in that column (left-aligned, space-padded)

### Allowed packages

The entire Go standard library.

### Expected type and function

```go
package challenge

type Boarding struct {
    Name string
    Seat string
    Gate string
}

func RenderBoardingLines(items []Boarding) []string {
    // solution
}
```

### Examples

Input:

```go
[]Boarding{
    {Name: " Ada ", Seat: " 12C ", Gate: " A1 "},
    {Name: "Bob", Seat: "", Gate: "A2"},
    {Name: "Cara", Seat: "7A", Gate: "B3"},
}
```

Expected:

```go
[]string{
    "Ada  | Seat 12C | Gate A1",
    "Cara | Seat 7A  | Gate B3",
}
```

The Name column is padded to 4 (longest valid name is "Cara"), the Seat column to 3 (longest valid seat is "12C"), the Gate column is not padded (only one width present after filtering). Bob is skipped, Seat is empty.

---

## Challenge 5c

### Instructions

An embedded flight-status display sends log messages over a constrained serial link. The logging library uses a minimal format string syntax to keep the binary small. It supports `%s` placeholders with optional width and alignment — nothing else.

Write `MiniPrintf` that formats a string using this minimal syntax.

### Supported format items

- literal characters (copied as-is)
- `%%` → literal percent sign `%` (consumes no argument)
- `%s` → consume one string argument
- `%<width>s` → right-align the argument in a field of `<width>` spaces (example: `%5s`)
- `%-<width>s` → left-align the argument in a field of `<width>` spaces (example: `%-5s`)

### Rules

- Width is optional; if present it must be an integer between 1 and 100 inclusive
- If the argument is already >= width characters, insert it with no padding, do not trim
- Each `%s` placeholder consumes exactly one argument
- The `-` flag requires a width. `%-s` with no digits is malformed
- Return `("", false)` if:
  - a placeholder is malformed (like `%`, `%-s`, `%0s`, `%5d`, `%--5s`, `%101s`)
  - there are not enough arguments
  - there are extra unused arguments
- Return `(result, true)` on success

### Allowed packages

The entire Go standard library except `fmt`.

### Expected function

```go
package challenge

func MiniPrintf(format string, args []string) (string, bool) {
    // solution
}
```

### Examples

```go
MiniPrintf("Gate %s ready", []string{"A1"})          // -> "Gate A1 ready", true
MiniPrintf("[%5s]", []string{"A1"})                  // -> "[   A1]", true
MiniPrintf("[%-5s]", []string{"A1"})                 // -> "[A1   ]", true
MiniPrintf("Load: 100%%", []string{})                // -> "Load: 100%", true
MiniPrintf("%s-%3s-%-4s", []string{"x", "y", "z"})   // -> "x-  y-z   ", true
MiniPrintf("%s-%s", []string{"HEL"})                 // -> "", false  // not enough arguments
MiniPrintf("%5d", []string{"12"})                    // -> "", false  // unsupported verb
MiniPrintf("%s", []string{"A1", "extra"})            // -> "", false  // extra unused argument
MiniPrintf("%101s", []string{"A"})                   // -> "", false  // width > 100
MiniPrintf("%-s", []string{"A1"})                    // -> "", false  // - without width
MiniPrintf("%0s", []string{"A1"})                    // -> "", false  // zero width
```

### Reference

The format string is scanned character by character. When a `%` is found, examine the following characters to determine what to emit:

- `%%`: emit a literal `%`, consume no argument
- `%s`: emit the next argument as-is
- `%<n>s` (positive integer 1–100): right-align — pad with spaces on the left so total width is `n`
- `%-<n>s` (dash then positive integer 1–100): left-align — pad with spaces on the right so total width is `n`
- anything else: malformed, return `("", false)`

Width padding example: `%5s` applied to `"Hi"` → `"   Hi"` (3 leading spaces, total length 5).