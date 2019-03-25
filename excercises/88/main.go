package main

import (
	"fmt"
)

// Failing Test cases - [],0,[2],1
func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	merge(arr1, 3, arr2, 3)
	fmt.Println(arr1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	for j < n && i < m+n {
		if nums1[i] > nums2[j] {
			InsertAt(nums1, i, nums2[j])
			j++
			i++
		} else {
			i++
		}
	}
	if j != n-1 || j == 0 {
		for j < n {
			InsertAt(nums1, m+j, nums2[j])
			j++
		}
	}
}

func InsertAt(arr []int, index int, val int) {
	prev := arr[index]
	arr[index] = val
	for i := index + 1; i < len(arr); i++ {
		prev, arr[i] = arr[i], prev
	}
}

//https://play.golang.org/p/n5i8oR0yyd8
