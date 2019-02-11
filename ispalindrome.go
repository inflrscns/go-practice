package main

import (
  "fmt"
  s "strings"
  re "regexp"
)

func reverse(str string) (new_str string) {
  for i := range str {
    new_str += string(str[len(str)-i-1])
  }
  return
}

func is_palindrome(str string) bool {
  str = s.ToLower(s.Replace(str, " ", "", -1))
  if match, _ := re.MatchString("[^a-zA-Z]+", str); !match {
    if len(str) > 1 && str == reverse(str) {
      return true
    }
  }
  return false
}

func main() {
  fmt.Println(reverse("hello"))
  fmt.Println(is_palindrome("test"))
  fmt.Println(is_palindrome("racecar"))
  fmt.Println(is_palindrome("hello"))
  fmt.Println(is_palindrome("a"))
  fmt.Println(is_palindrome(""))
  fmt.Println(is_palindrome("nurse i spy gypsies run"))
  fmt.Println(is_palindrome("a man a plan a canal Panama"))
  fmt.Println(is_palindrome("go hang a salami im a lasagna hog"))
}
