# Challenge 2

## Challenge 2a: Warm-up

### Instructions

This challenge verifies your understanding of integers and how they are represented in memory, as well as the basic arithmetic or bitwise operations.

Implement a function `CountBits` that takes in a positive number. Return how many bits are set in the in-memory representation of that number.

For instance, the number 2534 is written `1001_1110_1111` in binary (underscores added for clarity). In this number, 9 bits are set (ie. are equal to 1). Therefore, `CountBits(2534)` should return 9.

### Allowed packages

None.

### Expected function

```go
package challenge

func CountBits(n uint64) int {
    //
}
```

### Examples

```go
CountBits(0)    // --> 0
CountBits(4)    // --> 1
CountBits(10)   // --> 2
CountBits(2543) // --> 9
```

### Reference

Summary of the bitwise operators in Go (note that the numbers in this table are binary literals, but the operators operate on any ints).

| Operation          | Result       | Description   |
|--------------------|--------------|---------------|
| `0b0011 & 0b0101`  | `0b0001`     | Bitwise AND   |
| `0b0011 \| 0b0101` | `0b0111`     | Bitwise OR    |
| `0b0011 ^ 0b0101`  | `0b0110`     | Bitwise XOR   |
| `0b00110101 << 2`  | `0b11010100` | Left shift    |
| `0b00110101 >> 2`  | `0b00001101` | Right shift   |

---

## Challenge 2b: Challenge

### Instructions

The integer types available in Go have a fixed bit size, meaning they can each represent only a specific range of values. In this challenge, you will implement an arbitrary precision integer type that can contain any positive integer (limited only by the machine's memory), and perform basic arithmetic operations on it. These are often called "big ints".

One possible, relatively inefficient way is to simply represent a number's digits as elements in a slice, and to perform operations the same way it is done by hand: digit by digit, not forgetting the carry.

### Expected function signatures

Define a `BigInt` type, and the following arithmetic functions:

- `Add(a, b BigInt) BigInt`: returns the sum of arguments `a` and `b` (that is, the equivalent of `a + b`).
- `Multiply(a, b BigInt) BigInt`: multiply `a` and `b` (that is, the equivalent of `a * b`).

Additionally, implement the following functions for converting to and from other types:

- `FromInt(i uint64) BigInt`: convert a `uint64` to a `BigInt` of the same value.
- `ToDecimal(i BigInt) string`: convert the argument to its minimal decimal representation (that is, a string of digits from 0 to 9 that does not start with zeros, unless is it exactly `"0"`).
- `FromDecimal(s string) BigInt`: create a `BigInt` from a string containing its decimal representation. This function should ignore any character which is not a digit from 0 to 9. A string that contains no digits is equivalent to the string `"0"`.

### Allowed packages

None.

### Example

```go
a := FromInt(1)
b := FromDecimal("1")

for range 200 {
    a, b = b, Add(a, b)
}

println(ToDecimal(a)) // --> 453973694165307953197296969697410619233826
```

---

## Challenge 2c: In-depth

### Instructions

Many file formats are binary: ie. they directly contain the bytes of the data, as opposed to text-based where the file contains human-readable textual representations of the data, often in ASCII.

In this challenge, you will parse a simple image file format, similar to BMP or PNG. Write a function `LoadImage` that takes in a slice of bytes and returns the image, as defined by the struct below.

### Format description

**Important note:** all integers in this data format are encoded in little-endian, meaning: least significant byte first. For instance, the number 305441741, which in hexadecimal is `0x1234ABCD`, would be represented in the byte slice as the bytes: `0xCD, 0xAB, 0x34, 0x12`.

#### Header

This image file format begins with a header that contains some information about the size and content.

| Offset (bytes) | Size (bytes) | Field description                |
|----------------|--------------|----------------------------------|
| 0              | 4            | The string `"IMAG"`             |
| 4              | 1            | Bit depth                        |
| 5              | 2            | Image width                      |
| 7              | 2            | Image height                     |
| 9              | 1            | Orientation                      |
| 10             | 6            | Vendor identifier string         |

- **Bit depth:** indicates how many bits are used for a single pixel. Valid values are 8, 16 or 24. Values 8 and 16 indicates a grayscale picture, while 24 indicates a truecolor image (see data description below).
- **Image width and height** are respectively the number of column and rows of pixels in the image. The data must contain exactly `width * rows` pixels.
- **Orientation:** a value of 0 indicates the data is encoded row by row, while a value of 1 indicates that the data is encoded column by column.
- **Vendor identifier** is a free UTF8 string that programs manipulating these images can set. Its length must always be exactly 6 bytes.

Note that the header is, by definition, always 16 bytes long.

#### Pixel data

Immediately after the header comes the pixel data:

- If the orientation is row by row, the data will contain all rows from top to bottom, with pixels inside each row from left to right. If the orientation is column by column, the data will contain all columns from left to right, with pixels inside each column from top to bottom.
- If the bit depth indicates a grayscale picture, each pixel will be a single number indicating intensity (0 is black, and the maximum possible value is white).
- If the bit depth is truecolor, each pixel will be exactly three bytes representing red, green and blue levels, in this order. In each color, 0 is no intensity and the maximum value is full intensity. For instance, a fully red pixel would be represented by the following bytes: `255, 0, 0`. A magenta pixel at half intensity would be: `127, 0, 127`.

**Note:** to represent gray in truecolor, all three channels (red, green, blue) are set to the same intensity. For instance:

- white: `255, 255, 255`
- black: `0, 0, 0`
- light grey at about 75% intensity: `190, 190, 190`

### Allowed packages

None.

### Expected types and function

```go
type Pixel struct {
    Red, Green, Blue float32  // Intensity in each channel ranges from 0.0 to 1.0
}

type Image struct {
    Width, Height  uint16
    Vendor         string
    Pixels         [][]Pixel  // Pixels must be row by row regardless of original orientation
}

func LoadImage(bytes []byte) *Image {

}
```

The function must return `nil` if any error is encountered during parsing (invalid value in any header field or incorrect pixel data length).