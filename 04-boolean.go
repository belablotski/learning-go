package main

import "fmt"

func main() {
  var a = true
  var b = false
  fmt.Println(a && a)
  fmt.Println(a && b)
  fmt.Println(a || a)
  fmt.Println(a || b)
  fmt.Println(!a)
}