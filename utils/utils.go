package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func InSlice(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DisplaySlice(slice []int, idx ...int) {

	redPrint := color.New(color.FgRed, color.Bold)
	for index, element := range slice {
		if InSlice(idx, index) {
			redPrint.Printf("% 4d", element)
		} else {
			fmt.Printf("% 4d", element)
		}
	}
	fmt.Println("")
}
