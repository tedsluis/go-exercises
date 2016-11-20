// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
var u = flag.Bool("u", false, "convert to uppercase")
var l = flag.Bool("l", false, "convert to lowercase")

func main() {
	flag.Parse()
	text := strings.Join(flag.Args(), *sep)
	if *u == true {
		text = strings.ToUpper(text)
	}
	if *l == true {
		text = strings.ToLower(text)
	}
	fmt.Print(text)
	if !*n {
		fmt.Println()
	}
}
