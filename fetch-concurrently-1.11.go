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
	for _, url := range os.Args[1:] {
		for times := 0; times < 10; times++ {
			go fetch(url, ch, startcount, times) // start a goroutine
			startcount++;
		}
	}
	var sortindex [] string
	lines := make(map[string]string)
	for count := 1; count < startcount; count++ {
		result := strings.Split(<-ch,";") // receive from channel ch
		index := fmt.Sprintf("%-20s %4s %4s", result[3], result[1], result[0])
		sortindex = append(sortindex,index)
		lines[index]=fmt.Sprintf("%-20s order=%4s overall=%4s sec=%4s nbytes=%7s", result[3],result[4],result[0],result[1],result[2])
	}
	sort.Strings(sortindex)
	for _,index := range sortindex {
		fmt.Fprintf(os.Stderr,"-> %-30s -> %-60s \n", index,lines[index])
		//fmt.Println(">"+index+">"+lines[index])
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string, count int, times int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%4d;%.2f;%30s;%s;%4d", count, secs, err, url, times) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%4d;%.2f;%7d;%s;%4d", count, secs, nbytes, url, times)
}
