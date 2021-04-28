package main

import "fmt"

func fib(x int) int {
	j, k := 0, 1

	if x == 0 {
		return 0
	}

	for i := 1; i < x; i++ {
		j, k = k, j+k
	}

	return k
}

func main() {
	fmt.Println(fib(50))
}
