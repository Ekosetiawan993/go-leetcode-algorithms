package main

import (
	"fmt"
)

func insertSort(arr []int) []int {
	// copy the array
	arr2 := make([]int, len(arr))
	copy(arr2, arr)

	for i := 0; i < len(arr2); i++ {
		for j := i; j > 0; j-- {
			if arr2[j] < arr2[j-1] {
				temp := arr2[j]
				arr2[j] = arr2[j-1]
				arr2[j-1] = temp
			}
			// fmt.Println()
		}
		// fmt.Println(i)
	}
	return arr2
}

func main() {
	arr := []int{4, 3, 2, 2, 1, 6, 10, 7, 4}
	arrSorted := insertSort(arr)
	fmt.Println("before", arr)
	fmt.Println(arrSorted)
}
