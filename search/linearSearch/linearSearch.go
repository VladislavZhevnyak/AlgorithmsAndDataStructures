package main

import "fmt"

func main() {
	arr := []int{8, 5, 9, 0, 4, 1}
	item := 0
	idx := linearSearch(arr, item)
	fmt.Println(idx)
}

func linearSearch(arr []int, item int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return i
		}
	}
	return -1
}

