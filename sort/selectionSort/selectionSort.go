package main

import "fmt"

func main() {
	//O(n^2)
	arr := []int{12, 23, 2, 17, 11, 0, 4, 29, 45}
	selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
