package main

import "fmt"

func genFactory(start uint) func() uint {
  // closure is n + f
  n := start
  f := func () uint { n++; return n }
  return f
}

func main() {
  gen := genFactory(0)
  fmt.Println(gen())
  fmt.Println(gen())
  fmt.Println(gen())
}
