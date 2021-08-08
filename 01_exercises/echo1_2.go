// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[0:] {
		fmt.Println("index", idx, "value:", arg)
	}
}
