package main

import (
	"fmt"
	"time"
)

func main() {
	const gigasec = int32(1e9)
	start := time.Date(1959, 07, 19, 0, 0, 0, 0, time.UTC)

	added := start.Add(time.Second * time.Duration(gigasec))

	// fmt.Printf("%v gigasec = %v years %v days %v hours %v mins %v secs", gigasec)
	fmt.Println(added)

}
