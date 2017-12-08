package main

import "fmt"

func main() {
  for i := 0; i < 10; i++ {
    if i == 0 {
      fmt.Println(i, "This is Zero")
    } else if i % 2 == 0 {
      fmt.Println(i, "It's even")
    } else {
      fmt.Println(i, "It's odd")
    }
  }
}
