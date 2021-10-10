# Golang First Steps

**Book [The Go Programming Language](http://www.gopl.io/)**

### [Example folder](https://github.com/franlopezm/golang_first_steps/tree/main/examples)
Contains the implementation of the examples in the book. They are all organized by chapter.

### DD_exercises
Solutions to the exercises included in the book. **DD** indicates the chapter of the book to which the exercises belong.

> All exercises must be run from its folder.

- [01 Tutorial](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/01_exercises.md)
- [02 Program Structure](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/02_exercises.md)

---

## Commands

- Run `go run FILE_NAME`
- Compile `go build FILE_NAME`
- Create a module `go mod init MODULE_NAME`, MODULE_NAME must be unique or the same project name.
- Download project dependencies `go mod tidy`

## Package ftm
### Printf
  - `%d`                decimal integer
  - `%x`, `%o`, `%b`    integer in hexadecimal, octal, binary
  - `%f`, `%g`, `%e`    floating-point number
  - `%t`                boolean
  - `%c`                rune (Unicode code point)
  - `%s`                string
  - `%q`                quoted string
  - `%v`                any value in a natural format
  - `%T`                type of any value
  - `%%`                literal percent sign (no operand)

---

## Program structure
### Variable
#### Basic declarations
`var name type = expression`

- **Omit type**, `var name = expression`. Type is determinated by the initializer expression.
- **Omit expression**, `var name type`. Initial value is:
  - 0 for numbers.
  - false for booleans.
  - "" for strings.
  - nil for interfaces, reference types (slice, pointer, map, channel, function).
- **Declaration of multile variables**:
  - `var i, j, k int`
  - `var b, f, s = true, "Hello", 3.98`. bool, string, float64

#### Short declarations
`name := expression`

#### Pointers
A pointer values is the address of a variable. Example:

```go
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // 1
*p = 2          // equivalent to x = 2
fmt.Println(x)  // 2
```
