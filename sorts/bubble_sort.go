package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	numbers []string
)

func bubbleSort(slice []string) []string {
	swapped := false
	for i := 0; i < len(slice)-1; i++ {
		// len(slice)-1-i 後面已經排序過了，所以無需再判斷
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = strings.TrimSpace(slice[j+1]), strings.TrimSpace(slice[j])
			}
		}
		//fmt.Println(strings.Join(slice, ","))
		if swapped {
			break
		}
	}
	return slice
}

func init() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter numbers separated by a comma(,): ")
	text, _ := reader.ReadString('\n')
	numbers = strings.Split(text, ",")
}

func main() {
	fmt.Println("Origin numbers: ")
	fmt.Println(strings.Join(numbers, ","))

	fmt.Println("Sorted numbers: ")
	fmt.Println(strings.Join(bubbleSort(numbers), ","))
}
