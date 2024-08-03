package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9, 12, 15, 18, 22, 24}
	item := 7
	index := binarySearch(arr, 0, len(arr)-1, item)
	fmt.Printf("Index of %d is %d\n", item, index)
}

func binarySearch(arr []int, left int, right int, item int) int {
	if left > right {
		return -1
	}
	middle := (left + right) / 2
	if arr[middle] == item {
		return middle
	}
	if item > arr[middle] {
		return binarySearch(arr, middle+1, right, item)
	} else {
		return binarySearch(arr, left, middle-1, item)
	}
}
