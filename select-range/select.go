package main

import (
	"fmt"
)

func printChan(c1 chan bool, c2 chan bool) {
	var x1, x2 = 1, 1
	for {
		select {
		case <-c1:
			fmt.Println("X1 ran", x1, "times")
			x1++
		case c2 <- true:
			fmt.Println("X2 ran", x2, "times")
			x2++
		}
	}
}

func main() {
	c1 := make(chan bool)
	c2 := make(chan bool)

	go printChan(c1, c2)

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			c1 <- true
		} else {
			<-c2
		}
	}
}
