# Golang First Steps

**Book [The Go Programming Language](http://www.gopl.io/)**

### [Example folder](https://github.com/franlopezm/golang_first_steps/tree/main/examples)
Contains the implementation of the examples in the book. They are all organized by chapter.

### DD_exercises
Solutions to the exercises included in the book. **DD** indicates the chapter of the book to which the exercises belong.

> All exercises must be run from your folder.

#### Tutotial (01_exercises)
- **[1_1_echo](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_1_echo)**, prints on screen the arguments passed when calling the function.

  `go run main.go Hello 230 ups "Be careful"`

- **[1_1_echo](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_2_echo)**, prints on screen the arguments passed when calling the function, and its index.

  `go run main.go Hello 230 ups "Be careful"`

- **[1_3_echo](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_3_echo)**, benchmark running time between our potentially inefficient versions and the one that uses strings.Join.

  `go run main.go Hello 230 ups "Be careful"`

- **[1_4_dup](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_4_dup)**, find duplicate lines and print names of all files in which each duplicated line occurs.

  `go run main.go text1.txt text2.txt text3.txt text4.txt`

- **[1_5_animated_gif](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_5_animated_gif)**, Create a green gif.

  `go run main.go > out.gif`

- **[1_6_animated_gif](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_6_animated_gif)**, Produce images in multiple colors.

  `go run main.go > out.gif`

- **[1_7_fetch](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_7_fetch)**, Fetching a URL.

  `go run main.go https://wikipedia.org`

- **[1_8_fetch](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_8_fetch)**, Fetching a URL, and add prefix "https://".

  `go run main.go wikipedia.org`

- **[1_9_fetch](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_9_fetch)**, Fetching a URL and print HTTP status code.

  `go run main.go wikipedia.org`

- **[1_10_fetch_concurrently](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_10_fetch_concurrently)**, Fetching URLs concurrently and write results in a file.

  `go run main.go https://es.wikipedia.org https://golang.org https://es.wikipedia.org/wiki/Universo_cinematogr%C3%A1fico_de_Marvel`

- **[1_12_server_lissajous](https://github.com/franlopezm/golang_first_steps/tree/main/01_exercises/1_12_server_lissajous)**, The Lissajous server, read parameter values from the URL query.

  `go run main.go`
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