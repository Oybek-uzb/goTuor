package main

import (
	"fmt"
	"time"
)

func ping(ping string, ch chan string) {
	ch <- ping
	pong := <-ch
	fmt.Println(pong)
	time.Sleep(1000 * time.Millisecond)
}

func pong(pong string, ch chan string) {
	ping := <-ch
	ch <- pong
	fmt.Println(ping)
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	const PING string = "PING"
	const PONG string = "PONG"

	ch := make(chan string)

	for {
		go ping(PING, ch)
		pong(PONG, ch)
	}
}
