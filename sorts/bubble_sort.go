package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter numbers separated by a comma(,): ")
	text, _ := reader.ReadString('\n')
	sli := strings.Split(text, ",")

	fmt.Println("Origin numbers: ")
	fmt.Println(strings.Join(sli, ","))

	fmt.Println("Sorted numbers: ")
	fmt.Println(strings.Join(bubbleSort(sli), ","))
}
