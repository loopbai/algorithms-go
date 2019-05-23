package main

import (
	"fmt"
)

func bubbleSort(slice []int) []int {
	var swapped bool
	for i := 0; i < len(slice)-1; i++ {
		swapped = false
		// len(slice)-1-i 後面已經排序過了，所以無需再判斷
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
			}
		}
		//fmt.Println(slice)
		if swapped == false {
			break
		}
	}
	return slice
}

func main() {
	numbers := []int{4, 2, 323, 44, 67, 4, 6, 87, 3, 56, 7, 3}

	fmt.Println("Origin numbers: ")
	fmt.Println(numbers)

	fmt.Println("Sorted numbers: ")
	fmt.Println(bubbleSort(numbers))
}
