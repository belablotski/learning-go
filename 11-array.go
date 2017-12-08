package main

import "fmt"

func sum(arr [3]float64) float64 {
  var res float64 = 0.0
  for i := 0; i < len(arr); i++ {
    res += arr[i]
  }
  return res
}

func sum2(arr [3]float64) float64 {
  var res float64 = 0.0
  for _, x := range arr {
    res += x
  }
  return res
}

func avg(arr [3]float64) float64 {
  return sum(arr) / float64(len(arr))
}

func main() {
  var arr [3]float64
  arr[0] = 5.0
  arr[1] = 7.0
  arr[2] = 3.6
  fmt.Println("Avg:", avg(arr))

  arr2 := [3]float64 {1.0, 1.5, 2.3}
  fmt.Println("Sum v1 and v2:", sum(arr2), sum2(arr2))
  fmt.Println("Avg:", avg(arr2))
}
