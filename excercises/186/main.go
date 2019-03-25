package main

import "fmt"

func main() {

	input := []byte("Hello World")
	reverseWords(input)

	fmt.Println(string(input))
}

func reverseWords(str []byte) {
	var result []string
	las_pos := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			result = append(result, string(str[las_pos:i]))
			las_pos = i + 1
		}
	}
	result = append(result, string(str[las_pos:]))
	//str = make([]byte, len(str))
	fmt.Println(str)
	las_pos = 0
	for i := len(result) - 1; i >= 0; i-- {
		for j := 0; j < len(result[i]); j++ {
			str[las_pos] = []byte(result[i])[j]
			las_pos++
		}
		if i != 0 {
			str[las_pos] = []byte(" ")[0]
			las_pos++
		}
	}
	fmt.Println(string(str))
}
