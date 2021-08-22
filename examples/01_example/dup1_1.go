// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
// Answer to insert data
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	enter := "Enter a line of data:"
	for fmt.Println(enter); input.Scan(); fmt.Println(enter) {
		line := input.Text()

		if len(line) == 0 {
			break
		}

		counts[line]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
