// Fetch prints the content found at a URL.
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
		if  strings.HasPrefix(url, "http://") != true {
			url  = "http://"+url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error, Get(url): %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: reading %s: %v\n", url, err)
			os.Exit(1)
   		}
		resp.Body.Close()
		fmt.Fprintf(os.Stdout, "HTTP Status: %s \n", resp.Status)
	}
}
