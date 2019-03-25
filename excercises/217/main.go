package main

import (
	"fmt"
)

func main() {
	arr := []int{7, 1, 1, 3, 6, 4}
	fmt.Println(containsDuplicate(arr))
}

func containsDuplicate(nums []int) bool {
	output := make(map[int]int)
	for index, v := range nums {
		if _, ok := output[v]; ok {
			return true
		} else {
			output[v] = index
		}
	}
	return false
}

//https://play.golang.org/p/DL3Gjpmill3
