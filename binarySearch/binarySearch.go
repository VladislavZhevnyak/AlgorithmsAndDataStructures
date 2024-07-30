package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9, 12, 15, 18, 22, 24}
	item := -5
	index := binarySearch(arr, item)
	fmt.Printf("Index of %d is %d\n", item, index)
}

func binarySearch(arr []int, item int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		value := arr[mid]
		if item == value {
			return mid
		} else if value > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
