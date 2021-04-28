package main

import (
	"fmt"
)

func oddEvenSum(arr [6]int) [2]int {
	oddSum, evenSum := 0, 0
	length := len(arr)

	for i := 0; i < length; i++ {

		if i%2 == 1 {
			oddSum += arr[i]
		}

		if arr[i]%2 == 1 {
			evenSum += arr[i]
		}
	}

	res := [2]int{oddSum, evenSum}

	return res
}

func main() {
	var arr = [6]int{13, 10, 1, 9, 8, 7}

	fmt.Println(oddEvenSum(arr))
}
