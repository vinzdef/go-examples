package main

import (
	"fmt"
	"time"
)

func say(word string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word, i)
	}
}

func main() {
	go say("Routine 1")
	go say("Routine 2")

	go (func(msg string) {
		go say(msg)
	})("Anonymous Routine")

	// To hang main thread
	fmt.Scanln()
}
