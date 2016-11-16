// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
	"sort"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	startcount := 1
	prefix := 1
	for _, url := range os.Args[1:] {
		for times := 0; times < 10; times++ {
			go fetch(url, ch, startcount, times, prefix) // start a goroutine
			startcount++
		}
		prefix++
	}
	var sortindex [] string
	lines := make(map[string]string)
	for count := 1; count < startcount; count++ {
		result := strings.Split(<-ch,";") // receive from channel ch
		index := fmt.Sprintf("%-20s %4s %4s", result[3], result[1], result[0])
		sortindex = append(sortindex,index)
		lines[index]=fmt.Sprintf("%-25s order=%7s overall=%4s sec=%4s nbytes=%7s, message=%s", result[3],result[4],result[0],result[1],result[2],result[5])
	}
	sort.Strings(sortindex)
	for _,index := range sortindex {
		fmt.Fprintf(os.Stderr,"%-60s \n", lines[index])
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string, count int, times int, prefix int) {
	start := time.Now()
	resp, err := http.Get(url)
	secs := time.Since(start).Seconds()
	if err != nil {
		ch <- fmt.Sprintf("%4d;%.2f;%7d;%s;%2d/%-4d;%s", count, secs, 0, url, prefix, times, err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	secs = time.Since(start).Seconds()
	if err != nil {
		ch <- fmt.Sprintf("%4d;%.2f;%7d;%s;%2d/%-4d;%s", count, secs, 0, url, prefix, times, err) // send to channel ch
		return
	}
	secs = time.Since(start).Seconds()
	ch <- fmt.Sprintf("%4d;%.2f;%7d;%s;%2d/%-4d;%s", count, secs, nbytes, url, prefix, times, "ok")
}
