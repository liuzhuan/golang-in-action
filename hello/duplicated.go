package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	arr2 := [5]int{13, 42, 56, 12, 11}
	arr1 = arr2
	fmt.Printf("arr1: %v\narr2: %v\n", arr1, arr2)
	arr2[1] = 1024

	fmt.Println("after updating")
	fmt.Printf("arr1: %v\narr2: %v\n", arr1, arr2)
}

// Output:
// arr1: [13 42 56 12 11]
// arr2: [13 42 56 12 11]
// after updating
// arr1: [13 42 56 12 11]
// arr2: [13 1024 56 12 11]

// In summary, the duplicated arrays are two different object