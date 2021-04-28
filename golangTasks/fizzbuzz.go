package main

import (
	"fmt"
)

func fizzBuzz(x int) {
	res := ""

	for i := 1; i <= x; i++ {
		if i%15 == 0 {
			res = "FizzBuzz"
		} else if i%5 == 0 {
			res = "Buzz"
		} else if i%3 == 0 {
			res = "Fizz"
		} else {
			res = fmt.Sprint(i)
		}
		fmt.Print(res, " ")
	}
	fmt.Println()
}

func main() {

	fizzBuzz(40)

}
