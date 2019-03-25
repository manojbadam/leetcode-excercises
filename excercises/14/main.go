package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	result := []byte{}
	if len(strs) > 0 {
		num_strings := len(strs)
		least_length := len(strs[0])
		for i := 1; i < num_strings; i++ {
			if len(strs[i]) < least_length {
				least_length = len(strs[i])
			}
		}

		for i := 0; i < least_length; i++ {
			compare_value := strs[0][i]
			for j := 1; j < num_strings; j++ {
				if strs[j][i] != compare_value {
					return string(result)
				}
			}
			result = append(result, strs[0][i])
		}
	}
	return string(result)
}

//https://play.golang.org/p/mvMSym9cf1n
