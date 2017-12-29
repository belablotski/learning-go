// Array parameters are passed by value, slice parameters are passed by reference

package main

import (
	"fmt"
)

func f1(a [10]int) {
	a[0] = 111
}

func f2(a []int) {
	a[0] = 111
}


func main() {
	a := [10]int{1,2,3,4}
	b := a[:]
	
	fmt.Println(a)
	f1(a)
	fmt.Println(a)
	fmt.Println()

	fmt.Println(a)
	fmt.Println(b)
	f2(b)
	fmt.Println(a)
	fmt.Println(b)
}
