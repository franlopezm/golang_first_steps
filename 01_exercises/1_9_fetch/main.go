// Fetch prints the content found at each specified URL, and HTTP status code.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		hasPrefix := strings.HasPrefix(url, "https://")
		if !hasPrefix {
			url = "https://" + url
		}

		res, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		bytes, err := io.Copy(os.Stdout, res.Body)
		res.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("\n\nNumber of bytes: %d\nHTTP status: %v\nHTTP code: %v\nHTTP proto: %v\n", bytes, res.Status, res.StatusCode, res.Proto)
	}
}
