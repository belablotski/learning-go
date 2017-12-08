// Slice is a segment of an array, but unlike array, the length is allowed to change

/* var x []float64
   x := make([]float64, lenght, capacity)
   or
   a := [5]float64{1,2,3,4,5}
   x := a[start:end]    // "start" is inclusive, "end" is exclusive index; a[:] == [0:len(a)]

   append(slice, new_elem_1,...)
   copy(dst_slice, src_slice)
*/

package main

import "fmt"

func main() {
  s1 := []int{10, 20, 30, 40}
  s2 := append(s1, 50, 60, 70)
  fmt.Println("Slices s1 and s2:", s1, s2)

  s3 := make([]int, 2, 3)
  copy(s3, s2)
  fmt.Println("Slice s3:", s3)    // [10, 20]
}
