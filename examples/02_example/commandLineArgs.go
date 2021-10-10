// Echo 4 prints its command line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline") // Command line -n avoid newline
var sep = flag.String("s", " ", "separator")           // Command line -s specify a separator

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
