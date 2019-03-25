package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 1, 2, 3}
	fmt.Println(plusOne(arr1))
}

func plusOne(digits []int) []int {
	digits[len(digits)-1] += 1
	tail := 0
	if digits[len(digits)-1] >= 10 {
		tail = digits[len(digits)-1] / 10
		digits[len(digits)-1] = digits[len(digits)-1] % 10
		for i := len(digits) - 2; i >= 0; i-- {
			digits[i] += tail
			tail = digits[i] / 10
			digits[i] = digits[i] % 10
		}
	}
	if tail > 0 {
		digits = append([]int{tail}, digits...)
	}
	return digits
}
