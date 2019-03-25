package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length := removeDuplicates(arr)
	fmt.Println(arr[0:length])
}

func removeDuplicates(nums []int) int {
	if len(nums) > 0 {
		count := len(nums)
		for left_ptr := 1; left_ptr < count; left_ptr++ {
			if nums[left_ptr-1] == nums[left_ptr] {
				slideNums(nums, left_ptr)
				left_ptr--
				count--
			}
		}
		return count
	} else {
		return 0
	}
}

func slideNums(nums []int, arr_pos int) {
	for i := arr_pos; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	nums = nums[0 : len(nums)-1]
}

//https://play.golang.org/p/YlRdJ4FZSzw
