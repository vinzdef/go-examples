package main

import (
	"fmt"
)

func increase(c chan int) {
	for i := 0; i < 3; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go increase(c)
	for v := range c {
		fmt.Println(v)
	}
}
