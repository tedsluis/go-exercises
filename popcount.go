// Popcount, bit count

package main

import (
	"fmt"
	"os"
	"popcount"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseUint(arg, 0, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		y := popcount.PopCount(x)
		popcount.Display(x)
		fmt.Println(" ,x=", x, ", y=", y) 
	}
}
