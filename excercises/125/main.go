package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s_arr := []byte(s)
	f_ptr := 0
	l_ptr := len(s_arr) - 1
	for f_ptr < l_ptr {
		if !((s_arr[f_ptr] > 47 && s_arr[f_ptr] < 58) || (s_arr[f_ptr] > 96 && s_arr[f_ptr] < 123)) {
			f_ptr++
			continue
		}
		if !((s_arr[l_ptr] > 47 && s_arr[l_ptr] < 58) || (s_arr[l_ptr] > 96 && s_arr[l_ptr] < 123)) {
			l_ptr--
			continue
		}
		if s[f_ptr] == s[l_ptr] {
			l_ptr--
			f_ptr++
			continue
		} else {
			return false
		}
	}
	return true
}

//https://play.golang.org/p/kn7jThOGIV-
