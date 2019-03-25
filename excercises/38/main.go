package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		value := []int{1}
		for i := 2; i <= n; i++ {
			temp := []int{}
			count := 1
			for j := 0; j < len(value)-1; j++ {
				if value[j] == value[j+1] {
					count++
				} else {
					temp = append(temp, count)
					temp = append(temp, value[j])
					count = 1
				}
			}
			temp = append(temp, count)
			temp = append(temp, value[len(value)-1])
			value = temp
		}
		barr := []byte{}
		for _, v := range value {
			barr = strconv.AppendInt(barr, int64(v), 20)
		}
		return string(barr)
	}
}

//https://play.golang.org/p/CLf4INQYhyD
