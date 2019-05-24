package main

import (
	"fmt"

	"github.com/loopbai/algorithms-go/utils"
)

func bubbleSort(slice []int) []int {
	var swapped bool
	for i := 0; i < len(slice)-1; i++ {
		swapped = false
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
				utils.DisplaySlice(slice, j, j+1)
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
	numbers := []int{4, 323, 44, 67, 4, 6, 87, 56, 7, 3}

	fmt.Println("Origin numbers: ")
	utils.DisplaySlice(numbers, -1, -1)

	fmt.Println("Sorting numbers: ")
	bubbleSort(numbers)
}
