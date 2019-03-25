package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 2, 1}
	fmt.Println(singleNumber(arr))
}

func singleNumber(nums []int) int {
	output := make(map[int]bool)
	for _, v := range nums {
		if _, ok := output[v]; ok {
			output[v] = !output[v]
		} else {
			output[v] = true
		}
	}
	for k, v := range output {
		if v {
			return k
		}
	}
	return 0
}

// https://play.golang.org/p/Xh7xahpboQV
