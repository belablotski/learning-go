package main

import "fmt"

type V struct {
  x, y int
}

func main() {
  m := make(map[string]V)
  m["test"] = V {1, 2}
  fmt.Printf("%+v\n", m)

  //m["test"].x = 5		// Error: "cannot assign to struct field m["test"].x in map"

  v1 := m["test"]		// Pull struct, make changes, push it back
  v1.x, v1.y = v1.y, v1.x
  m["test"] = v1
  fmt.Printf("%+v\n", m)

  v2 := m["test2"]		// V with x = y = 0
  fmt.Printf("%+v\n", v2)
  fmt.Printf("%+v\n", m)
  m["test2"] = v2
  fmt.Printf("%+v\n", m)
}
