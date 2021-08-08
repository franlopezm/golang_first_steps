// Echo program benchmark running time between our potentially inefficient versions and the one that uses strings.Join
package main

import (
	"os"
	"strings"
	"time"
)

func main() {
	limit := 10_000_000

	timeTest1 := bench(limit, test1)
	timeTest2 := bench(limit, test2)
	timeTest3 := bench(limit, test3)

	println("test 1", timeTest1)
	println("test 2", timeTest2)
	println("test 3", timeTest3)
}

// Run tests
func bench(limit int, toRun func()) int64 {
	start := time.Now()
	for i := 0; i < limit; i++ {
		toRun()
	}
	return time.Since(start).Milliseconds()
}

// First echo
func test1() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
}

// Echo with range
func test2() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
}

// Echo efficient
func test3() {
	strings.Join(os.Args[1:], " ")
}
