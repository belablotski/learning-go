package main

import "fmt"

func stats(numbers []float64) (min, max, sum, avg float64) {
  min = numbers[0]
  max = numbers[0]
  sum = 0.0
  for _, x := range numbers {
    if x < min { min = x }
    if x > max { max = x }
    sum += x
  }
  avg = sum / float64(len(numbers))
  return
  // = return min, max, sum, avg
}

// Variadic parameters
func stats2(numbers ...float64) (min, max, sum, avg float64) {
  return stats(numbers)
}

func main() {

  gen := func () []float64 {
    return []float64{10.0, 20.0, 30.0, 40.0}
  }

  nums := gen()
  min, max, sum, avg := stats(nums)
  fmt.Println(min, max, sum, avg)
  min, max, sum, avg = stats2(nums...)    // Pass slice as variadic parameter
  fmt.Println(min, max, sum, avg)
}
