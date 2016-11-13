package main

import "fmt"

func main () {
  a:=1
  for a <=100 {
    b:=1
    for b <=100 {
      x = (a - b)
      fmt.Print(a,"*",b,"=",a * b,"  ")
      b+=1
    }
    a+=1
  }
}
