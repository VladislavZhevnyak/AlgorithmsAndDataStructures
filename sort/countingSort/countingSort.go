package main

import "fmt"

func main() {
	//O(n)
	arr := []int{3, 2, 7, 0, 1, 4}
	countingSort(arr)
	fmt.Println(arr)
}

func countingSort(arr []int) {
	k := findMaxValue(arr)
	count := make([]int, k+1)

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	fmt.Println(count)
	result := make([]int, 0, len(arr))
	for i := 0; i < len(count); i++ {
		for count[i] > 0 {
			result = append(result, i)
			count[i]--
		}
	}
	copy(arr, result)
}

func findMaxValue(arr []int) int {
	maxValue := arr[0]
	for _, val := range arr {
		if val >= maxValue {
			maxValue = val
		}
	}
	return maxValue
}

