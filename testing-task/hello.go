package main

import (
	"fmt"
)

func hello(str string) string {

	if len(str) == 0 {
		return fmt.Sprint("Hello, Dude")
	} else {
		return fmt.Sprintf("Hello, %v", str)
	}
}
