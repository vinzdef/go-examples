package main

import "visibility"

func main() {
	// Ok!
	visibility.VisibleVar = 9001
	flag := visibility.VisibleMethod(true)

	// Panic
	visibility.anotherVar = 9002
}
