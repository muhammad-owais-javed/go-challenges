# Challenge 6

## Challenge 6a

### Instructions

Evaluate nested arithmetic expressions. Expressions are represented as nested arrays.

Each expression is either:

- an integer, or
- an operation array

Operations have the form: `[operator, left, right]`

### Supported operators

- `add`: addition
- `sub`: subtraction
- `mul`: multiplication

Return the evaluated integer result.

The input will always follow the specified format and you may assume all values and operators are valid. You do not have to handle over- or underflows.

### Allowed packages

None.

### Expected function and type

```go
package challenge

func Evaluate(expr any) int {
    // solution
}
```

### Examples

```go
expr := []any{
    "sub",
    9,
    6,
}

Evaluate(expr)
// -> 3
```

```go
expr := []any{
    "sub",
    []any{"mul", 2, 5},
    []any{"add", 10, 8},
}

Evaluate(expr)
// -> -8
```

```go
expr := []any{
    "add",
    99,
    []any{
        "add",
        3,
        []any{
            "mul",
            7,
            7,
        },
    },
}

Evaluate(expr)
// -> 151
```

### Reference

`i, ok := v.(int)` performs a type assertion; `ok` indicates whether `v` is of type `int`, and `i` holds the value if it is (otherwise zero value).

---

## Challenge 6b

### Instructions

You are given a simplified HTML DOM tree. Each node represents an HTML element and may contain child elements.

Write a function that finds all elements with a given tag name and returns their full paths from the root element separated by ` > `.

Some nodes may have an empty tag (`""`). These represent anonymous wrapper elements used for layout — they exist in the tree but carry no semantic meaning. Empty-tag nodes should still be traversed, but should not appear in the generated paths.

Finding a match does not stop the search. If a node matches, keep searching its children — a descendant might also match.

The search tag will never be empty (`""`).

### Return value

- Return a slice of path strings, one per match (e.g. `"html > body > p"`).
- Return paths in depth-first order (top to bottom).
- If root is `nil` or there are no matches, return `nil`.

Return paths in the order they are found, from top to bottom. This is called depth-first traversal.

### Allowed packages

`strings`

### Expected function and type

```go
package challenge

type Node struct {
    Tag      string
    Children []*Node
}

func FindPaths(root *Node, tag string) []string {
    // solution
}
```

### Examples

Simple case:

```go
root := &Node{
    Tag: "html",
    Children: []*Node{
        {
            Tag: "body",
            Children: []*Node{
                {Tag: "p"},
            },
        },
    },
}

FindPaths(root, "p")
// → []string{"html > body > p"}
```

With anonymous wrapper nodes and multiple matches:

```go
root := &Node{
    Tag: "html",
    Children: []*Node{
        {
            Tag: "body",
            Children: []*Node{
                {
                    Tag: "",
                    Children: []*Node{
                        {Tag: "p"},
                        {Tag: "span"},
                    },
                },
                {
                    Tag: "section",
                    Children: []*Node{
                        {
                            Tag: "div",
                            Children: []*Node{
                                {Tag: "p"},
                            },
                        },
                    },
                },
                {
                    Tag: "p",
                    Children: []*Node{
                        {
                            Tag: "div",
                            Children: []*Node{
                                {Tag: "p"},
                            },
                        },
                    },
                },
            },
        },
    },
}

FindPaths(root, "p")
// → []string{
//     "html > body > p",
//     "html > body > section > div > p",
//     "html > body > p",
//     "html > body > p > div > p",
// }
```

### Reference

Depth-first search (DFS) explores one path as far as possible before backtracking, while breadth-first search (BFS) explores all neighboring nodes at the current depth before moving deeper.

---

## Challenge 6c

### Instructions

In this task, you will implement a parser and evaluator for simple S-expressions. The simple syntax of S-expressions made them popular as the basis for a variety of programming and configuration languages, notably the Lisp family.

In this example, they are used to represent basic arithmetic and logic.

### Syntax

An expression is one of the following:

- An integer literal: an optional leading `-`, followed by one or more decimal digits. Examples: `0`, `42`, `-7`.
- A compound expression: `( operator arg... )`.

Whitespace (spaces, tabs, newlines) is insignificant, except to separate any two elements of a compound expression.

### Operators

| Operator | Syntax         | Description                              |
|----------|----------------|------------------------------------------|
| add      | `(add A B)`    | A + B                                    |
| sub      | `(sub A B)`    | A − B                                    |
| mul      | `(mul A B)`    | A × B                                    |
| if       | `(if C T E)`   | If C ≠ 0, return T; otherwise return E   |

### Conditionals

`(if C T E)` evaluates `C` first. If the result is nonzero, evaluates and returns `T`. Otherwise, evaluates and returns `E`.

### Error cases

Return `(0, false)` if:

- an operator receives the wrong number of arguments
- an unrecognized operator is used
- parentheses are unmatched or misplaced
- there is trailing content after the top-level expression
- any token is otherwise invalid (e.g. `--1`, `1abc`, `abc?def`, `+`, etc.)

You can use `int` as the type for all numerical values, and you do not need to handle integer over- or underflows.

### Allowed packages

The entire Go standard library

### Expected function

```go
package challenge

func EvaluateSExpression(expr string) (int, bool) {
    // solution
}
```

### Examples

Arithmetic:

```go
EvaluateSExpression("42")                        // -> 42, true
EvaluateSExpression("-7")                        // -> -7, true
EvaluateSExpression("(add 3 (mul 4 5))")         // -> 23, true
EvaluateSExpression("(sub(add 10 2)(mul 2 3))")  // -> 6,  true  (whitespace optional around parens)
EvaluateSExpression("(add 1)")                   // -> 0,  false (wrong arity)
EvaluateSExpression("(div 4 2)")                 // -> 0,  false (unknown operator)
EvaluateSExpression("(add 1 2) 3")               // -> 0,  false (trailing content)
```

Conditionals:

```go
EvaluateSExpression("(if 1 42 0)")                   // -> 42, true
EvaluateSExpression("(if 0 42 99)")                  // -> 99, true
EvaluateSExpression("(if (sub 5 5) 42 99)")          // -> 99, true
EvaluateSExpression("(if 1 42 (add 999 999))")       // -> 42, true
EvaluateSExpression("(if (mul 2 3) 15 (sub 1 1))")   // -> 15, true
```