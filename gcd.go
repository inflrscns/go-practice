package main

import (
  "fmt"
)

func gcd(a,b int) int {
  if b == 0 {
    return a
  } else {
    return gcd(b, a%b)
  }
}

func main() {
  fmt.Println(gcd(180, 120))
}
