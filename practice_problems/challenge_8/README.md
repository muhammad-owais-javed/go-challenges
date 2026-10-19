# Challenge 8

## Challenge 8a

### Instructions

A telemetry pipeline receives streams of single-digit sensor codes (0-9). Before writing a stream to disk, consecutive repeated codes are compressed using run-length encoding, to save storage space.

Write a function `EncodeRLE` that encodes a string of digits into a compact, comma-separated form.

### Rules

- Output is a comma-separated list of `"count:digit"` pairs, in the order the runs appear in the input.
- Runs are maximal: all consecutive occurrences of the same digit are grouped into a single pair.
- Empty input produces an empty string (not `"0:"`).
- The input is guaranteed to contain only ASCII digits (0-9).

### Allowed packages

The entire Go standard library.

### Expected function

```go
package challenge

func EncodeRLE(s string) string {
    // solution
}
```

### Examples

```go
EncodeRLE("")        // -> ""
EncodeRLE("5")       // -> "1:5"
EncodeRLE("111")     // -> "3:1"
EncodeRLE("1112233") // -> "3:1,2:2,2:3"
EncodeRLE("121212")  // -> "1:1,1:2,1:1,1:2,1:1,1:2"
```

### Reference

A telemetry pipeline is the infrastructure that collects data emitted by remote devices or services, such as metrics, logs, or sensor readings, and carries it through stages of transport, compression, storage, and analysis. Compressing repetitive data close to the source, before it is written or sent onward, is a common early stage, since it cuts down on bandwidth and storage costs. Run-length encoding is one of the simplest ways to do this: it replaces runs of identical values with a single count, which works well when the same reading repeats many times in a row.

---

## Challenge 8b

### Instructions

Every record exported by the platform is assigned a numeric ID. When a record is referenced in a URL, its ID is encoded as a short, URL-safe string using a base-62 alphabet (digits, then uppercase letters, then lowercase letters) — similar to how link-shortening services or video platforms generate compact IDs.

Write functions to encode a `uint64` into base 62, and to decode it back.

### Format description

The base-62 alphabet, indexed from 0 to 61:

```
0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
```

(10 digits, then 26 uppercase letters, then 26 lowercase letters.)

### Rules

- `EncodeBase62` never produces a string with a leading `'0'`, except for the representation of 0 itself, which is `"0"` (not `""`).
- `DecodeBase62` must accept any string made only of alphabet characters, including ones with redundant leading `'0'`s (they don't affect the numeric value, the same way `"007"` and `"7"` mean the same thing in decimal).
- `DecodeBase62` must return an error if the input is empty, or if it contains any character outside the alphabet. In both cases, the returned `uint64` must be `0`.
- The alphabet is case-sensitive: `'A'` (value 10) and `'a'` (value 36) are different digits.
- You don't need to handle overflow: every string given to `DecodeBase62` in tests represents a value that fits in a `uint64`.

### Reference

Converting a number to an arbitrary base is a generalization of converting decimal to binary or hexadecimal: repeatedly divide by the base, and collect the remainders. Decoding does the reverse: it reads the string as a positional numeral system, the same way `"123"` in decimal means `1*100 + 2*10 + 3*1`.

### Allowed packages

`errors`.

### Expected functions

```go
package challenge

func EncodeBase62(n uint64) string {
    // solution
}

func DecodeBase62(s string) (uint64, error) {
    // solution
}
```

### Examples

```go
EncodeBase62(0)      // -> "0"
EncodeBase62(9)      // -> "9"
EncodeBase62(10)     // -> "A"
EncodeBase62(35)     // -> "Z"
EncodeBase62(36)     // -> "a"
EncodeBase62(61)     // -> "z"
EncodeBase62(62)     // -> "10"
EncodeBase62(123456) // -> "W7E"

DecodeBase62("W7E")  // -> 123456, nil
DecodeBase62("00A")  // -> 10, nil
DecodeBase62("")     // -> 0, error
DecodeBase62("W7E!") // -> 0, error
```

---

## Challenge 8c

### Instructions

The platform exports records to partner systems using a compact binary format: type-length-value (TLV). Write functions to encode a list of records into this format, and to decode them back.

### Format description

Each record is encoded as three parts, back to back:

| Field   | Size (bytes) | Description                 |
|---------|--------------|-----------------------------|
| Type    | 1            | The record's type, as given |
| Length  | 2            | Payload length, big-endian  |
| Payload | Length       | The raw payload bytes       |

"Big-endian" means the most significant byte of a multi-byte number comes first. For instance, the length 2 is encoded as the two bytes `0x00, 0x02`.

Multiple records are encoded by concatenating their TLV sequences one after another, with nothing in between.

Finally, the entire byte sequence is converted into a lowercase hexadecimal string, where every byte is represented by exactly 2 hex characters.

### Reference

Each byte splits into two 4-bit halves (nibbles), each represented by one hex character:

| Value    | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 |
|----------|---|---|---|---|---|---|---|---|---|---|----|----|----|----|----|----|
| Hex char | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | a  | b  | c  | d  | e  | f  |

For instance, the byte `0x4B` (75 in decimal) is represented as the two characters `"4b"`.

### Worked example

`Record{Type: 1, Payload: []byte("hi")}` encodes as follows:

- Type byte: `01`
- Length (2 bytes, big-endian): `0002` (`len("hi")` is 2)
- Payload bytes, hex-encoded: `6869` (ASCII `'h'` is `0x68`, `'i'` is `0x69`)
- Full encoding: `"0100026869"`

Two records concatenate directly:

`Record{Type: 1, Payload: []byte("hi")}`, `Record{Type: 2, Payload: []byte("!")}` → `"010002686902000121"`

### Rules

- `EncodeRecords` returns an error if any record's payload is longer than 65535 bytes (the length field only has 2 bytes to work with).
- An empty (or nil) records slice encodes to an empty string, with no error.
- `DecodeRecords` returns an error if the input string has odd length, or contains characters outside `0-9a-fA-F`.
- `DecodeRecords` returns an error if the decoded bytes are truncated: an incomplete header (fewer than 3 bytes remaining where a new record should start), or a payload shorter than its declared length.
- An empty input string decodes to an empty (or nil) slice of records, with no error.

### Allowed packages

`errors`.

### Expected types and functions

```go
package challenge

type Record struct {
    Type    byte
    Payload []byte
}

func EncodeRecords(records []Record) (string, error) {
    // solution
}

func DecodeRecords(s string) ([]Record, error) {
    // solution
}
```

### Examples

```go
EncodeRecords([]Record{{Type: 1, Payload: []byte("hi")}})
// -> "0100026869", nil

EncodeRecords([]Record{
    {Type: 1, Payload: []byte("hi")},
    {Type: 2, Payload: []byte("!")},
})
// -> "010002686902000121", nil

EncodeRecords([]Record{{Type: 0xAB, Payload: []byte{0xCA, 0xFE, 0xBA, 0xBE}}})
// -> "ab0004cafebabe", nil

EncodeRecords([]Record{{Type: 5, Payload: bytes.Repeat([]byte{0xAB}, 300)}})
// -> "05012c" + strings.Repeat("ab", 300), nil

DecodeRecords("0100026869")
// -> []Record{{Type: 1, Payload: []byte("hi")}}, nil

DecodeRecords("ab0004cafebabe")
// -> []Record{{Type: 0xAB, Payload: []byte{0xCA, 0xFE, 0xBA, 0xBE}}}, nil

DecodeRecords("010002")
// -> nil, error

DecodeRecords("0100026")
// -> nil, error
```