package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		temp := []int{}
		result := true
		for x > 0 {
			temp = append(temp, x%10)
			x = x / 10
		}

		for i := 0; i < len(temp)/2; i++ {
			if temp[i] == temp[len(temp)-i-1] {
				continue
			} else {
				result = false
				break
			}
		}
		return result
	}
}

//https://play.golang.org/p/7VvpXe7CSXH
