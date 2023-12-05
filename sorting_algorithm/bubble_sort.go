package main

import (
	"fmt"
)

func BubbleSort(arr []int) []int {
	arr2 := make([]int, len(arr))
	copy(arr2, arr)

	for i, _ := range arr2 {
		isSorted := true
		for j := 0; j < (len(arr) - i - 1); j++ {

			if arr2[j] > arr2[j+1] {
				temp := arr2[j]
				arr2[j] = arr2[j+1]
				arr2[j+1] = temp
				isSorted = false
			}
		}
		fmt.Println("sort iter:", i)
		// if the array already sorted
		if isSorted {
			return arr2
		}
	}
	return arr2
}

func main() {
	// arr := []int{3, 1, 2, 3, 5, 2, 2, 7, 9, 6, -1}
	arr := []int{2, 1, 3, 4, 5}
	fmt.Println("Before ordering:", arr)

	sortedArr := BubbleSort(arr)
	fmt.Println(sortedArr)

}
