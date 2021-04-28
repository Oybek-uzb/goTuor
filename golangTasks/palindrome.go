package main

import (
	"fmt"
	"strings"
)

func palindrome(text string) bool {

	arr := strings.Split(text, "")
	length := len(arr)
	isSame := true

	for i := 0; i < length/2; i++ {
		if arr[i] != arr[length-i-1] {
			isSame = false
			break
		}
	}

	return isSame
}

func main() {
	fmt.Println(palindrome("abcba"))
}
