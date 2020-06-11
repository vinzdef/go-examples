package main

import "fmt"

func main() {
	i := 0
	fmt.Println(i)

	for i = 1; i < 4; i++ {
		defer fmt.Println(i)
	}
}
