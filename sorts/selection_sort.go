package main

import (
	"fmt"

	"github.com/loopbai/algorithms-go/utils"
)

func selectionSort(slice []int) []int {
	length := len(slice)
	minIndex := 0
	for i := 0; i < length; i++ {
		minIndex = i
		for j := i; j < length; j++ {
			if slice[minIndex] > slice[j] {
				minIndex = j
			}
		}
		if i != minIndex {
			slice[minIndex], slice[i] = slice[i], slice[minIndex]
			utils.DisplaySlice(slice, minIndex, i)
		}
	}

	return slice
}

func main() {
	numbers := []int{4, 2, 323, 44, 67, 4, 6, 87, 3, 56, 7, 3}

	fmt.Println("Origin numbers: ")
	utils.DisplaySlice(numbers, -1, -1)

	fmt.Println("Sorting numbers: ")
	selectionSort(numbers)
}
