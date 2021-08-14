// Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]

	counts := make(map[string]int)
	fileNames := make(map[string][]string)

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames, os.Args[0])
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts, fileNames, arg)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			files := strings.Join(fileNames[line], "\n")
			fmt.Printf("%v\n\t%d\t%s\n", files, n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string, filename string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		line := input.Text()

		counts[line]++
		if !contains(fileNames[line], filename) {
			fileNames[line] = append(fileNames[line], filename)
		}
	}
}

func contains(a []string, t string) bool {
	for _, arg := range a {
		if arg == t {
			return true
		}
	}

	return false
}
