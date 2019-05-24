package main

import (
	"fmt"

	"github.com/loopbai/algorithms-go/utils"
)

func insertionSort(slice []int) []int {

	for out := 1; out < len(slice); out++ {
		var movedIdx []int
		temp := slice[out]
		in := out

		for ; in > 0 && slice[in-1] >= temp; in-- {
			slice[in] = slice[in-1]
			movedIdx = append(movedIdx, in)
		}
		slice[in] = temp

		movedIdx = append(movedIdx, in)
		utils.DisplaySlice(slice, movedIdx...)
	}

	return slice
}

func main() {
	numbers := []int{4, 323, 44, 67, 4, 6, 87, 56, 7, 3}

	fmt.Println("Origin numbers: ")
	utils.DisplaySlice(numbers)

	fmt.Println("Sorting numbers: ")
	insertionSort(numbers)
}
