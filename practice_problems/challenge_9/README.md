# Challenge 9

## Challenge 9a

### Instructions

A pipeline deduplicates a stream of string tags before forwarding them downstream. Write a function `Deduplicate` that removes repeated tags while preserving order.

The standard library's `slices.Compact` only removes consecutive duplicates, and expects the input already sorted — it can't dedupe an arbitrary-order stream like this one. That's why this is worth hand-rolling.

### Rules

- Return a new slice containing each element of `items` exactly once: only the first occurrence of each value is kept, in the order it first appears.
- The input slice must not be modified by the call.
- Empty input returns an empty slice (nil or zero-length are both fine).

### Allowed packages

The entire Go standard library.

### Expected function

```go
package challenge

func Deduplicate(items []string) []string {
    // solution
}
```

### Examples

```go
Deduplicate([]string{"a", "b", "a"})                // -> []string{"a", "b"}
Deduplicate([]string{"x", "y", "z"})                // -> []string{"x", "y", "z"}
Deduplicate([]string{"a", "a", "a"})                // -> []string{"a"}
Deduplicate([]string{"c", "a", "b", "a", "c", "b"}) // -> []string{"c", "a", "b"}
Deduplicate([]string{})                             // -> []string{}
```

---

## Challenge 9b

### Instructions

A build system runs a test matrix. The matrix is described as a list of groups. For instance, one group of operating systems, one of architectures, one of compiler versions, and so on. The system has to run one build for every possible way of picking a single option out of each group.

Write a function `Expand` that produces that list of builds.

### Rules

- Every returned row picks exactly one option from each group, in group order: `row[i]` is one of the options in `groups[i]`, and every row has exactly `len(groups)` options in it.
- Every possible row appears exactly once. The rows may come out in any order at all — no particular ordering is required.
- If any group is empty, return an empty result (nil or a zero-length slice are both fine).
- The returned rows must be independent: changing an element of one row must never change another row, and must never change any of the groups.
- `groups` and the option slices inside it must not be modified by the call.

### Allowed packages

None.

### Expected function

```go
package challenge

func Expand(groups [][]string) [][]string {
    // solution
}
```

### Examples

The rows are shown below in a convenient order, but any order is accepted — what matters is that every row is there, and that none of them is there twice.

```go
Expand([][]string{
    {"linux", "darwin"},
    {"amd64", "arm64"},
})
// -> [][]string{
//     {"linux", "amd64"},
//     {"linux", "arm64"},
//     {"darwin", "amd64"},
//     {"darwin", "arm64"},
// }
```

```go
Expand([][]string{
    {"linux"},
    {"amd64", "arm64"},
    {"1.24", "1.25"},
})
// -> [][]string{
//     {"linux", "amd64", "1.24"},
//     {"linux", "amd64", "1.25"},
//     {"linux", "arm64", "1.24"},
//     {"linux", "arm64", "1.25"},
// }
```

```go
Expand([][]string{
    {"linux", "darwin"},
    {"amd64"},
    {"1.25"},
    {"race", "norace"},
})
// -> [][]string{
//     {"linux", "amd64", "1.25", "race"},
//     {"linux", "amd64", "1.25", "norace"},
//     {"darwin", "amd64", "1.25", "race"},
//     {"darwin", "amd64", "1.25", "norace"},
// }
```

```go
Expand([][]string{{"linux", "darwin", "windows"}})
// -> [][]string{{"linux"}, {"darwin"}, {"windows"}}
```

```go
Expand([][]string{
    {"linux", "darwin"},
    {},
    {"1.25"},
})
// -> [][]string{}   (or nil)
```

```go
Expand([][]string{})
// -> [][]string{{}}
```

---

## Challenge 9c

### Instructions

A debugger inspects a running program that carves slices out of a few large shared buffers. It is handed two of those slices and nothing else — no buffer, no indices, no record of where either one came from. It has to work out whether the two are views on the same memory, and if they are, exactly which part of each one.

This is not a question about values. Two slices can hold identical values while being views on completely separate memory, and vice versa.

Write a function `Overlap` that answers this question.

### Rules

- If some part of `a` and some part of `b` are views on the same memory, return `ai`, `bi` and `n` such that `a[ai:ai+n]` and `b[bi:bi+n]` are views on exactly that shared memory: the same `n` elements, of the same array. `n` must be as large as possible.
- If no part of `a` is a view on the same memory as any part of `b`, return `0, 0, 0`. This covers two separate situations: the two slices may be views on entirely different arrays, or they may be views on the same array without sharing any of its memory.
- Everything must be left exactly as it was found: when `Overlap` returns, `a`, `b` and the arrays they are views on must hold the values they held when it was called.

### Allowed packages

None.

### Expected function

```go
package challenge

func Overlap(a, b []int) (ai, bi, n int) {
    // solution
}
```

### Examples

All the examples below use the same two buffers:

```go
arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
other := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
```

```go
Overlap(arr[2:5], arr[4:8])    // -> 2, 0, 1   (both are views on arr[4])
Overlap(arr[1:9], arr[3:6])    // -> 2, 0, 3
Overlap(arr[0:3], arr[3:6])    // -> 0, 0, 0
Overlap(arr[0:2], arr[5:7])    // -> 0, 0, 0
Overlap(arr[2:5], other[2:5])  // -> 0, 0, 0
Overlap(arr[0:0], arr[0:5])    // -> 0, 0, 0

Overlap(arr[2:5], append([]int(nil), arr[2:5]...))   // -> 0, 0, 0
```

### Reference

A slice value in Go is not the data itself. It is three things: a pointer to the array holding the data, a length — how many elements the slice is a view on — and a capacity, which counts the elements available in that array from where the slice begins to the end of the array. Length and capacity are independent, and a slice may be resliced beyond its length, up to its capacity.

Two slices can be views on the same array. When they are, a write through one of them is visible through the other, wherever the two are views on the same memory.