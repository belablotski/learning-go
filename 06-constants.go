package main

import "fmt"

func main() {
  const x string = "abc"
  const y = "def"
  fmt.Println(x + y)

  const (
    a = 1
    b = 2
    c = 3
  )
  fmt.Println(a, b, c)
}
