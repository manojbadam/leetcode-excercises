package main

import (
	"fmt"
)

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	rotate(arr, 3)
	fmt.Println(arr)
}

func rotate(nums []int, k int) {
	count := len(nums)
	for i := 0; i < k; i++ {
		temp := nums[count-1]
		slideNums(nums, 0)
		nums[0] = temp
	}
}

func slideNums(nums []int, arr_pos int) {
	for i := len(nums) - 2; i >= arr_pos; i-- {
		nums[i+1] = nums[i]
	}
}

//https://play.golang.org/p/wOpe986WAul
