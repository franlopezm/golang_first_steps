// FetchAll fetches URLs in parallel and reports their times, sizes and cache data.
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Start a goroutine
	}

	f, err := os.OpenFile("./logFile.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("Openfile Error %v\n", err)
		os.Exit(1)
	}

	w := bufio.NewWriter(f)
	for range os.Args[1:] {
		fmt.Fprintf(w, "%v\n", <-ch) // receive from channel ch
	}

	fmt.Fprintf(w, "%.2fs elapsed\n", time.Since(start).Seconds())
	w.Flush()
	f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // send to channel
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nBytes, url)
}
