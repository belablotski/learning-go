package main

import "fmt"

func main() {
  var digits map[string]int
  fmt.Println(digits["one"])    // prints 0 - the "empty" value of type int

  //digits["one"] = 1     -- map ins't initialized

  digits = make(map[string]int)
  digits["one"] = 1
  digits["two"] = 2
  digits["three"] = 3
  fmt.Println(digits["one"])
  fmt.Println(digits["two"])
  fmt.Println(digits["three"])

  if four, ok := digits["four"]; ok {
    fmt.Println(four)
  } else {
    fmt.Println("There is no element [four]")
  }

  digits["four"] = 4

  if four, ok := digits["four"]; ok {
    fmt.Println(four)
  } else {
    fmt.Println("There is no element [four]")
  }

  delete(digits, "four")

  if four, ok := digits["four"]; ok {
    fmt.Println(four)
  } else {
    fmt.Println("There is no element [four]")
  }
}
