package main

import (
	"fmt"

	"github.com/loopbai/algorithms-go/utils"
)

func quickSort(slice []int) []int {
	length := len(slice)
	if length < 2 {
		return slice
	}

	var left, right, sorted []int
	pivot := slice[0]
	for i := 1; i < length; i++ {
		switch {
		case slice[i] < pivot:
			left = append(left, slice[i])
		case slice[i] >= pivot:
			right = append(right, slice[i])
		default:
			fmt.Println("Error...")
		}
	}

	sorted = append(sorted, quickSort(left)...)
	sorted = append(sorted, pivot)
	sorted = append(sorted, quickSort(right)...)
	utils.DisplaySlice(sorted)
	return sorted
}

func main() {
	numbers := []int{4, 323, 44, 67, 4, 6, 87, 56, 7, 3}

	fmt.Println("Origin numbers: ")
	utils.DisplaySlice(numbers, -1, -1)

	fmt.Println("Sorting numbers: ")
	quickSort(numbers)
}
