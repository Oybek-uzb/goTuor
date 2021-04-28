package main

import (
	"fmt"
)

func cutDuplicates(arr []int) []int {
	slc := make([]bool, 0)
	res := make([]int, 0)

	length := len(arr)

	for i := 0; i < length; i++ {
		isExist := false
		for j := 0; j < length; j++ {
			if i != j && arr[i] == arr[j] {
				isExist = true
				break
			}
		}
		slc = append(slc, isExist)
	}

	for i := 0; i < length; i++ {
		if !slc[i] {
			res = append(res, arr[i])
		}
	}

	return res
}

func main() {
	arr := []int{1, 2, 1, 3, 4, 5, 8, 1, 9, 5, 5, 3}
	nonDupArr := cutDuplicates(arr)

	fmt.Println(nonDupArr)
}
