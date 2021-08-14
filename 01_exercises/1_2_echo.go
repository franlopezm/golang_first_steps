// Echo program print index and value, one per line
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
