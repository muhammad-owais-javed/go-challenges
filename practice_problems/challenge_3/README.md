# Challenge 3

## Challenge 3a

### Instructions

In this challenge, you will convert a string containing a date from one format to another. Specifically, from an ISO date and 24h time to a format closer to North American usage:

- **Input format:** `YYYY-MM-DD HH:MM` (all numerical, padding zeroes, 24h format)
- **Output format:** `D Month YYYY, H:MM` followed by `"AM"` or `"PM"` (no padding zeros for the day, full month name in letters, time in 12h format)

If parsing fails, return empty string.

### Allowed packages

The entire Go standard library. The `time` package may be particularly useful.

### Expected function

```go
package challenge

func FormatSessionTime(input string) string {
    // solution
}
```

### Examples

```go
FormatSessionTime("2024-02-28 19:32")
// "28 February 2024, 7:32PM"

FormatSessionTime("2001-10-03 03:55")
// "3 October 2001, 3:55AM"

FormatSessionTime("March 10 1998 10AM")
// ""
```

---

## Challenge 3b

### Instructions

In this challenge, you will build a conflict index for sessions coming from devices in different timezones. Each session contains an ID, a local Start, a local End, and a Timezone.

Return a conflict map where:

- **key:** session ID
- **value:** sorted list of conflicting session IDs

### Overlap rule

Two sessions overlap if one starts while the other one is ongoing.

Boundary examples (intervals are half-open):

- `[10,20)` and `[20,30)` do not overlap
- `[10,20)` and `[19,30)` do overlap
- `[10,20)` and `[10,20)` do overlap

### Rules

- assume all session IDs are unique and non-empty
- parse local datetimes as `YYYY-MM-DD HH:MM` (24 hours format)
- timezone offsets use format `±HH:MM` (for example: `+02:00`, `-05:30`)
- if a session has invalid datetime/offset, or start >= end after conversion, ignore it
- include only sessions with at least one conflict
- no duplicates in value lists
- don't include a session ID in its own conflict list
- for each key, sort conflict IDs ascending lexicographically

If parsing fails for a session, treat that session as invalid and ignore it.

### Allowed packages

The entire Go standard library. The `time` package is expected.

### Expected function

```go
package challenge

type ZonedSession struct {
    ID       string
    Start    string // local datetime
    End      string // local datetime
    Timezone string // "±HH:MM"
}

func BuildConflictMap(sessions []ZonedSession) map[string][]string {
    // solution
}
```

---

## Challenge 3c

### Instructions

You are working on an embedded scheduler prototype where the Go `time` package is unavailable. Dates are represented as simple `"YYYY-MM-DD"` strings. During disruption handling, dates need to be shifted by fixed number of days. You will implement your own logic for that purpose.

Write a function `AddDays(date string, days int) (string, bool)` which takes in a date in the format indicated above, and a number of days (possibly negative). It must return a date in the same format, where `days` calendar days have been added (or subtracted if negative).

It should return a second parameter indicating if the input string was valid and if the resulting date is within the valid range (see below).

### Rules

- input date must be exactly `"YYYY-MM-DD"` (4 digits, `-`, 2 digits, `-`, 2 digits)
- Year, month, day must form a valid Gregorian calendar date. If they don't, return an empty string and `false`.
- If the resulting year would be negative or over 9999, return an empty string and `false`.
- On success, return the result in the same `"YYYY-MM-DD"` format, and `true`.

**Note:** in the Gregorian calendar, a year is a leap year if:

- it is divisible by 400
- or: it is divisible by 4 but not by 100

### Allowed packages

None.

### Expected function

```go
package challenge

func AddDays(date string, days int) (string, bool) {
    // solution
}
```

### Sample input/output

```go
AddDays("2024-02-28", 1)  // -> "2024-02-29", true
AddDays("2023-02-28", 1)  // -> "2023-03-01", true
AddDays("1900-02-28", 1)  // -> "1900-03-01", true
AddDays("2024-02-29", -1) // -> "2024-02-28", true
AddDays("2020-01-01", -1) // -> "2019-12-31", true
AddDays("2024-13-01", 10) // -> "", false
```

### Reference

| Month     | Month Number | Days                  |
|-----------|--------------|-----------------------|
| January   | 1            | 31                    |
| February  | 2            | 28/29 on a leap year  |
| March     | 3            | 31                    |
| April     | 4            | 30                    |
| May       | 5            | 31                    |
| June      | 6            | 30                    |
| July      | 7            | 31                    |
| August    | 8            | 31                    |
| September | 9            | 30                    |
| October   | 10           | 31                    |
| November  | 11           | 30                    |
| December  | 12           | 31                    |