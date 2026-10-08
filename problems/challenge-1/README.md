# Go Coding Challenges

## Challenge 1a: ParsePassengerLine

### Instructions

Write a `ParsePassengerLine` function that normalizes raw string input into clean `Passenger` values (and report which lines are invalid).

The format of a single input line is as follows: `booking_id|name|seat`. Your goal is to accept valid rows and reject bad ones.

An input line is considered valid if and only if:

- it contains exactly 3 fields separated by `|`
- fields are trimmed of whitespace
- all fields are non-empty after trim

If the input is invalid, return a zero-value struct and `false`.

### Allowed Packages

The entire Go standard library.

### Expected Type and Function

```go
package challenge

type Passenger struct {
    BookingID string
    Name      string
    Seat      string
}

func ParsePassengerLine(line string) (Passenger, bool) {
    // solution
}
```

### Examples

```go
ParsePassengerLine("BK-123456|Ada Lovelace|12C")
// -> Passenger{BookingID: "BK-123456", Name: "Ada Lovelace", Seat: "12C"}, true

ParsePassengerLine("BK-123456| |12C")
// -> Passenger{}, false   // empty name after trim

ParsePassengerLine("BK-123456|Ada Lovelace")
// -> Passenger{}, false   // not exactly 3 fields
```

---

## Challenge 1b: ParsePassengerManifest

### Instructions

Write a `ParsePassengerManifest` function that takes an entire passenger manifest of passenger information, one record per line, and normalizes each valid string input into clean `Passenger` values. The batch may contain noise and duplicates. Keep valid passengers, and report which non-empty lines are invalid.

### Rules

- split input on `\n`
- for each line, trim spaces and trim trailing `\r` (Windows `\r\n` support)
- ignore empty lines
- report invalid non-empty lines
- if same `booking_id` repeats, keep the last valid one
- return all parsed records in a map indexed by booking ID, and a slice of invalid line numbers (1-based, ascending order)

An input line is considered valid if and only if:

- it contains exactly 3 fields separated by `|`
- fields are trimmed of whitespace before validation and returned trimmed
- all fields are non-empty after trim
- seat format is `<row><letter>` where:
  - row is a positive integer (>= 1) (but not bigger than max int)
  - letter is a capital letter, one of A to F
  - example valid seat: `12C`
  - `0A` is invalid
  - `1A` is valid
  - `001A` is valid

### Allowed Packages

The entire Go standard library.

### Expected Type and Function

```go
package challenge

type Passenger struct {
    BookingID string
    Name      string
    Seat      string
}

func ParsePassengerManifest(block string) (map[string]Passenger, []int) {
    // solution
}
```

### Sample Input / Output

**Input block:**

```
BK-123456|Ada Lovelace|12C
BK-000001|Ada Lovelace|12Z
BK-654321|Grace Hopper|7A
BK-123456|Ada L.|14F
```

**Expected:**

```go
map[string]Passenger{
    "BK-123456": {BookingID: "BK-123456", Name: "Ada L.", Seat: "14F"},
    "BK-654321": {BookingID: "BK-654321", Name: "Grace Hopper", Seat: "7A"},
}, []int{2}
```

---

## Challenge 1c: DecodeSeatMap

### Format Description

Seat occupancy comes from a compact device format.
You receive one line per row in this format:

```
<row_number>:<run_length_encoding>
```

Where run-length encoding is a sequence of:

```
<count><status>
```

**Status symbols:**

- `E` = empty
- `O` = occupied

**Example:**

`12:3E2O1E` means row 12 has seats: `E E E O O E`

### Task

Parse all valid lines into a map and add the line number of any invalid lines into a slice.

Decode each row into a boolean slice where:

- `false` = empty seat (E)
- `true` = occupied seat (O)

Return:

- map `row_number -> []bool`
- invalid line numbers (1-based, ascending)

### Validation Rules

- line must contain exactly one `:`
- row number must be an integer >= 1 (but not bigger than max int)
- run-length tokens must be well-formed (count then E or O)
- count must be integer >= 1 (but not bigger than max int)
- leading zeros in counts are allowed (e.g. `01E` is valid)
- empty run-length part is invalid
- ignore empty lines after trimming
- trim spaces around the whole line before parsing
- spaces inside the encoded payload are not allowed (e.g. `12:3E 2O` is invalid)
- duplicate row numbers:
  - any invalid occurrences are reported in invalid line numbers
  - among valid occurrences, the last valid occurrence wins (overwrite in the output map)

### Allowed Packages

The entire Go standard library.

### Expected Function

```go
package challenge

func DecodeSeatMap(block string) (map[int][]bool, []int) {
    // solution
}
```

### Examples

**Input block:**

```
12:3E2O1E
bad line
7:1O1E2O
3:0E2O
```

**Expected:**

```go
map[int][]bool{
    12: {false, false, false, true, true, false},
    7:  {true, false, true, true},
}, []int{2, 4}
```
