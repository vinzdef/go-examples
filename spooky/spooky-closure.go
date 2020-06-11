package main

import "fmt"

func intGenerator(start int) func() int {
	return func() int {
		defer func() { start++ }()
		return start
	}
}

func main() {
	next := intGenerator(0)

	for i := 0; i < 10; i++ {
		fmt.Print(next())
	}
}
