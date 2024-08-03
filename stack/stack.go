package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if isValid(s) {
		fmt.Println("String is valid")
	} else {
		fmt.Println("String isn't valid")
	}
}

func isValid(s string) bool {
	stack := []rune{}
	checkSymb := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != checkSymb[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}

