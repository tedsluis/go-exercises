// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	//"io/ioutil"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error, Get(url): %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: reading %s: %v\n", url, err)
			os.Exit(1)
   		}
		resp.Body.Close()
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s", b)
	}
}
