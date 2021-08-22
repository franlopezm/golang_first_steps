# First Steps With Golang

## Book [The Go Programming Language](http://www.gopl.io/)

### Example folder
Contains the implementation of the examples in the book. They are all organized by chapter.

### DD_exercises
Solutions to the exercises included in the book. **DD** indicates the chapter of the book to which the exercises belong.

> All exercises must be run from your folder.

#### Tutotial (01_exercises)
- **[1_1_echo]()**, prints on screen the arguments passed when calling the function.
  `go run main.go Hello 230 ups "Be careful"`
- **[1_1_echo]()**, prints on screen the arguments passed when calling the function, and its index.
  `go run main.go Hello 230 ups "Be careful"`
- **[1_3_echo]()**, benchmark running time between our potentially inefficient versions and the one that uses strings.Join.
  `go run main.go Hello 230 ups "Be careful"`
- **[1_4_dup]()**, find duplicate lines and print names of all files in which each duplicated line occurs.
  `go run main.go text1.txt text2.txt text3.txt text4.txt`

---

## Commands

- Run `go run FILE_NAME`
- Compile `go build FILE_NAME`

## package ftm
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