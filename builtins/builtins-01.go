package main

import "fmt"

func main() {
	slice := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		slice = append(slice, i)
	}

	fmt.Println(len(slice)) // 5
	fmt.Println(cap(slice)) // 6 (grew once)

	channel := make(chan bool, 3)
	channel <- true
	fmt.Println(len(channel)) // 1
	fmt.Println(cap(channel)) // 3

	keyValue := make(map[string]int, 3)
	keyValue["hello"] = 1
	fmt.Println(len(keyValue)) // 0
}
