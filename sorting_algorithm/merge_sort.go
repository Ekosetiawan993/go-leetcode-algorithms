package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	leftArr := mergeSort(arr[:mid])
	rightArr := mergeSort(arr[mid:])

	return merge(leftArr, rightArr)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// merge leftover elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	fmt.Println(result)

	return result
}

func main() {
	arr := []int{2, 6, 2, 1, 5, 6, 7, 4, 3, 9, 1}
	// arr := []int{0}
	arr2 := make([]int, len(arr))
	copy(arr2, arr)
	arrSorted := mergeSort(arr2)
	fmt.Println("Before sorting", arr)
	fmt.Println("After sorting", arrSorted)
	// fmt.Println(arr[1:])
}
