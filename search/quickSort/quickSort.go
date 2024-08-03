package main

import "fmt"

func main() {
	arr := []int{3, -5, 0, -6, 9, 29, 3, 25}
	arr = quickSort(arr)
	fmt.Println(arr)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int
	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}
