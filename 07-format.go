package main

import "fmt"

func main() {
  var (
    name string
    age int8
  )
  fmt.Print("Name: ")
  fmt.Scanf("%s\n", &name)
  fmt.Print("Age: ")
  fmt.Scanf("%d", &age)
  fmt.Printf(`Person %s
             is %d years old`, name, age)
}
