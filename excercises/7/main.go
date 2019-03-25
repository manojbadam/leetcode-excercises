package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	neg_num := false
	if x < 0 {
		neg_num = true
		x = 0 - x
	}
	temp_arry := []int{}
	for x > 0 {
		temp_arry = append(temp_arry, x%10)
		x = x / 10
	}
	sum := 0
	for _, v := range temp_arry {
		sum += v
		sum *= 10
	}
	sum = sum / 10
	if sum > math.MaxInt32 {
		return 0
	}
	if neg_num {
		return 0 - sum
	} else {
		return sum
	}
}
