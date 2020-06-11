package main

import "fmt"

func describe(unk interface{}) {
	fmt.Printf("Value is of type: %T \n", unk)
}
func main() {
	describe(42)
	describe("Hello")
}
