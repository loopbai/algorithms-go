package main

import (
	"fmt"
)

func quickSort(slice []int) []int {

	length := len(slice)
	if length < 2 {
		return slice
	}

	var left, right, qSorted []int
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
	qSorted = append(qSorted, quickSort(left)...)
	qSorted = append(qSorted, pivot)
	qSorted = append(qSorted, quickSort(right)...)
	return qSorted
}

func main() {
	numbers := []int{4, 2, 323, 44, 67, 4, 6, 87, 3, 56, 7, 3}

	fmt.Println("Origin numbers: ")
	fmt.Println(numbers)

	fmt.Println("Sorted numbers: ")
	fmt.Println(quickSort(numbers))
}
