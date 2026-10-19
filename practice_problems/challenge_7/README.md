# Challenge 7

## Challenge 7a

### Instructions

Write a function `RowWithMostEmptySpaces` that returns the index of the row containing the greatest number of empty parking spaces.

A parking lot is represented by a two-dimensional slice of runes:

- `'E'` = Empty parking space
- `'O'` = Occupied parking space

If multiple rows contain the maximum number of empty spaces, return the index of the first such row.

If all spaces are occupied or there are no empty parking spaces in the slice, return `-1`.

### Additional constraints

- The parking lot contains always at least one row.
- Each row may contain zero or more parking spaces.
- Each row can be of different length.
- Only the characters `'E'` and `'O'` will be present.

### Allowed packages

The entire Go standard library.

### Expected function and type

```go
package challenge

func RowWithMostEmptySpaces(parking [][]rune) int {
    // solution
}
```

### Examples

```go
RowWithMostEmptySpaces([][]rune{
    {'E', 'O', 'E', 'E'},
    {'O', 'O', 'E', 'O'},
    {'E', 'E', 'E', 'O'},
}) // -> 0

RowWithMostEmptySpaces([][]rune{
    {'E', 'O', 'O'},
    {},
    {'O', 'E', 'O'},
    {'E', 'E'},
}) // -> 3

RowWithMostEmptySpaces([][]rune{
    {'O', 'O', 'O'},
    {'O', 'O', 'O'},
    {'O', 'O', 'O'},
}) // -> -1
```

---

## Challenge 7b

### Instructions

Write a function `IsValidParkingLot` that determines whether a parking lot map is valid.

The parking lot is represented as a string. Each line represents a row of the map.

The following symbols may appear in the map:

- `'E'` = Empty parking space
- `'O'` = Occupied parking space
- `'W'` = Wall
- `'G'` = Gate

A parking lot is considered valid if all of the following conditions are met:

- The map contains at least one gate (`'G'`). Gates can only appear on the edge of the map.
- The map contains only valid characters and newline characters (`'\n'`).
- Every cell on the outer border of the map is either a wall (`'W'`) or a gate (`'G'`).
- Any parking space (`'E'` or `'O'`) must be completely surrounded (up, down, left and right) by other parking spaces, a wall or a gate.
- Rows may have different lengths.
- The parking lot has at least one parking space (empty or occupied).

### Additional constraints

The map string never starts or ends with a newline character, and never contains two consecutive newline characters (there are no empty lines).

### Allowed packages

The entire Go standard library.

### Expected function and type

```go
package challenge

func IsValidParkingLot(parking string) bool {
    // solution
}
```

### Examples

The maps below are laid out across multiple lines for readability. As you should know, a regular `"..."` string in Go can't contain a raw newline - build these strings with escaped `\n` instead (e.g `"WWGW\nWEOW\nWWWW"`).

```go
IsValidParkingLot(
    "WWGW\n" +
    "WEOW\n" +
    "WWWW")
// -> true

IsValidParkingLot(
    "WWWWW\n" +
    "WEOEW\n" +
    "WWWWW")
// -> false (no gate)

IsValidParkingLot(
    "WWGW\n" +
    "WEOE\n" +
    "WWWW")
// -> false (parking space on the border)

IsValidParkingLot(
    "WWGWW\n" +
    "WEOW\n" +
    "WWWWW")
// -> true

IsValidParkingLot(
    "WWGW\n" +
    "WEO\n" +
    "WWWW")
// -> false (the occupied space touches the outside)

IsValidParkingLot(
    "WWGW\n" +
    "WWWW")
// -> false (no parking space)
```

---

## Challenge 7c

### Instructions

Write a function `LargestEmptySquare` that finds the largest square region made up entirely of empty cells (`'E'`) in a 2D grid, and returns its side length and the row and column of its top-left corner.

The grid is represented as a 2D slice of runes containing only:

- `'E'` = Empty cell
- `'O'` = Obstacle

### Additional constraints

- If multiple squares share the same maximum side length, return the one whose top-left corner appears first when scanning top to bottom, left to right.
- If no empty cell exists, return a side length of `0` and coordinates `0, 0`.
- The grid contains at least one row and each row contains at least one cell.
- All rows have the same length.

### Allowed packages

The entire Go standard library.

### Expected function and type

```go
package challenge

func LargestEmptySquare(grid [][]rune) (size, row, col int) {
    // solution
}
```

### Examples

```go
LargestEmptySquare([][]rune{
    {'E', 'E', 'E'},
    {'E', 'E', 'E'},
    {'E', 'E', 'E'},
})
// -> size=3, row=0, col=0

LargestEmptySquare([][]rune{
    {'E', 'E', 'O'},
    {'E', 'E', 'E'},
    {'O', 'E', 'E'},
})
// -> size=2, row=0, col=0

LargestEmptySquare([][]rune{
    {'O', 'E', 'E'},
    {'O', 'E', 'E'},
    {'O', 'E', 'E'},
})
// -> size=2, row=0, col=1

LargestEmptySquare([][]rune{
    {'O', 'O', 'O'},
    {'O', 'O', 'O'},
})
// -> size=0, row=0, col=0

LargestEmptySquare([][]rune{
    {'E', 'O', 'E', 'E'},
    {'E', 'O', 'E', 'E'},
    {'E', 'E', 'E', 'E'},
})
// -> size=2, row=0, col=2 (two 2x2 squares exist: top-right and bottom; top-right comes first)
```