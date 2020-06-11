package main

import (
	"fmt"
	"time"
)

func read(flag chan bool) {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Read", <-flag)
}

func main() {
	flag := make(chan bool)

	go read(flag)
	fmt.Println("Waiting...")
	flag <- true

	// Thread locks
	// until `flag` can be written
	fmt.Println("Unlocked")
}
