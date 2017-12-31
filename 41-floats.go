package main

import (
  "fmt"

  "gonum.org/v1/gonum/floats"
)

func main() {
  fmt.Println(floats.Round(3.1415, 0))
  fmt.Println(floats.Round(3.1415, 1))
  fmt.Println(floats.Round(3.1415, 2))

  fmt.Println(floats.Round(3.5152, 0))
  fmt.Println(floats.Round(3.5152, 1))
  fmt.Println(floats.Round(3.5152, 2))
}