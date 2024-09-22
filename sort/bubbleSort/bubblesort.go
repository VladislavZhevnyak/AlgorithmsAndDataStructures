package main

import "fmt"

func bubbleSort(arr []int) {
	//O(n^2)
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 5}
	fmt.Println("Unsorted array:", arr)

	bubbleSort(arr)

	fmt.Println("Sorted array:", arr)
}
