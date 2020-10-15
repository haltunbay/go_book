package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	urls := os.Args[1:]
	start := time.Now()

	ch := make(chan string)

	for _, url := range urls {
		go fetch3(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("Total took: %.4fs\n", secs)
}

func fetch3(url string, ch chan string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("fetch error %v", err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("error while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	result := fmt.Sprintf("%.4fs  %7d %s status: %s", secs, nbytes, url, resp.Status)
	ch <- result
}

func fetch2(url string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %v\n", err)
		os.Exit(-1)
	}
	fmt.Println("response status:", resp.Status)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		fmt.Printf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	result := fmt.Sprintf("%.4fs  %7d %s", secs, nbytes, url)
	fmt.Println(result)

}

func fetch1(url string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %v\n", err)
		os.Exit(-1)
	}
	fmt.Println("response status:", resp.Status)
	io.Copy(os.Stdout, resp.Body)

}
