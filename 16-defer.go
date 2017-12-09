// defer = "finally"

package main

import (
  "fmt"
  "math/rand"
  "time"
  "os"
)

func demo1() {
  defer fmt.Println("Demo end")
  fmt.Println("Demo begin")
  if rand.Float64() > 0.5 {
    fmt.Println("Demo v1")
    return
  } else {
    fmt.Println("Demo v2")
    return
  }
  fmt.Println("Dear code")
}

func demo2() {
  fmt.Println("0")
  defer fmt.Println("7")
  defer fmt.Println("6")
  defer fmt.Println("5")
  fmt.Println("1")
  fmt.Println("2")
  fmt.Println("3")
  fmt.Println("4")
}

func demo3() {
  const filename = "16-defer.txt"
  f, _ := os.Create(filename)
  defer f.Close()
  f.WriteString("test")
}

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  demo1()
  demo2()
  demo3()
}
