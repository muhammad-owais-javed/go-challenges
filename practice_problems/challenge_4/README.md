# Challenge 4

## Challenge 4a

### Instructions

Write a function `SortHouses` that sorts a slice of houses from least valuable to most valuable, according to the following criteria:

1. More surface is more valuable
2. At equal surfaces, a house with more bedrooms is more valuable
3. At equal surfaces and bedrooms, a house with a garden is more valuable
4. All other things being equal, a house closer to the metro station is more valuable

### Allowed packages

The entire Go standard library.

### Expected type and function

```go
package challenge

type House struct {
    Surface         float64
    Bedrooms        uint
    HasAGarden      bool
    DistanceToMetro float64
}

func SortHouses(houses []House) {
    // solution
}
```

### Reference

In the Go documentation you might see a prototype similar to this:

```go
func FunctionName[S ~[]E, E any](x S, funcA func(a, b E) int)
```

Short explanation of `[S ~[]E, E any]`:

These are type parameters (generics). You can read this as:

- `E any` → the `E` type can be any type.
- `S ~[]E` → `S` must be a type whose underlying type is `[]E` (a slice of `E`).

The `~` means: not necessarily exactly `[]E`, but any named type based on `[]E`.

---

## Challenge 4b

### Instructions

You have been using the standard library to sort slices. Now you will implement a sorting algorithm yourself.

Implement `ComputeStats`, which takes a slice of players and returns:

- all players tied for first place (highest score), in no particular order
- the median score
- the average score

Computing the median will require sorting the scores first. You may implement any sorting algorithm you like. Your code will be tested with relatively short slices, so it doesn't have to be very efficient.

Return `nil, 0, 0` for an empty list.

### Allowed packages

None.

### Expected types and function

```go
package challenge

type Player struct {
    Name  string
    Score int
}

func ComputeStats(players []Player) (winners []string, median, average float64) {
    // solution
}
```

### Examples

```go
ComputeStats([]Player{
    {Name: "Charlie", Score: 42},
    {Name: "Alice", Score: 100},
    {Name: "Bob", Score: 42},
    {Name: "Diana", Score: 75},
})
// -> []string{"Alice"}, 58.5, 64.75

ComputeStats([]Player{
    {Name: "Charlie", Score: 100},
    {Name: "Alice", Score: 100},
    {Name: "Bob", Score: 42},
})
// -> []string{"Alice", "Charlie"}, 100, 80.666
```

### Reference

The median of a list of numbers is its middle value: there are as many values above and below. The easiest way to compute the median is to first sort the list of numbers, which puts the median at center of the array. For a list with an even number of values, the median is the average of the two middle values.

- `[90, 10, 20]` → median is 20
- `[30, 100, 10, 20]` → median is (20 + 30) / 2 = 25.0

The average of a list of numbers is the sum of all values divided by the count of values.

- `[90, 10, 20]` → average is (90 + 10 + 20) / 3 = 40.0
- `[30, 100, 10, 20]` → average is (30 + 100 + 10 + 20) / 4 = 40.0

---

## Challenge 4c

### Instructions

You have implemented sorting yourself, but have you wondered about efficiency?

In this challenge, you will be tasked with writing a sorting algorithm, but there will be an efficiency constraint: you cannot use simple sorts like bubble sort or selection sort.

Your function will act on an interface type that exposes the number of elements to sort, a way to compare two elements, and the possibility to swap two elements. The tester will track how many comparisons your algorithm performs, and will fail if it is over 5 times more than a reference implementation.

Your algorithm will be tested with random data, in order to avoid edge cases causing poor efficiency.

### Allowed packages

None.

### Expected type and function

```go
package challenge

type SortableContainer interface {
    Length() int          // Returns the number of elements in the container
    Less(i, j int) bool   // Returns true if the i-th element is strictly smaller than the j-th element
    Swap(i, j int)        // Swaps the i-th and j-th elements
}

func EfficientSort(container SortableContainer) {
    // solution here
}
```

### Reference

Quick-sort is one of the most famous efficient sorting algorithms. On average, it performs O(n log n) comparisons to sort n items, and it can be applied in place.

It works as follows:

1. Pick a pivot in the array, for instance the last element.
2. Rearrange the array so that: all values smaller than the pivot come first, then the pivot, then all the values larger than the pivot.
3. Recursively sort the low values, and recursively sort the high values.