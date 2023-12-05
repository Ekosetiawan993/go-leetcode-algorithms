package main

import (
	"fmt"
)

func selectionSort(arr []int) []int {
	arr2 := make([]int, len(arr))
	copy(arr2, arr)

	// iteration for all of the elements
	for i, n := range arr2 {
		currentMin := n
		currentMinPos := i
		// iteration for searching the smallest number
		for j := i; j < len(arr2); j++ {
			if arr2[j] < currentMin {
				currentMin = arr2[j]
				currentMinPos = j
			}
		}
		temp := n
		arr2[i] = currentMin
		arr2[currentMinPos] = temp

	}

	return arr2
}

func main() {
	arr := []int{2, 6, 2, 1, 5, 6, 7, 4, 3, 9, 1}
	// arr := []int{0}
	arrSorted := selectionSort(arr)
	fmt.Println("Before sorting", arr)
	fmt.Println("After sorting", arrSorted)
}
