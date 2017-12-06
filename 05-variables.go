package main;

import "fmt";

func main() {
  var x, y int32;
  x = 1
  y = 2
  fmt.Println(x + y);

  var s1 string = "abc"
  //fmt.Println(s + x + y);  -- invalid operation: s + x (mismatched types string and int32)

  s2 := "def"   // var s2 = "def"
  fmt.Println(s1 + s2)

  var (
    a float32 = 1.234
    b float64 = 4.321
  )
  fmt.Println(a, b)
}
