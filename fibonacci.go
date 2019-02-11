package main

import (
  "fmt"
)

func fib(n int) {
  a, b := 0, 1
  for i:=0; i<n; i++ {
    fmt.Println(a)
    a, b = b, a+b
  }
}

func main() {
  fib(13)
}
