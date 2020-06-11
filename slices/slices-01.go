package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	s := a[:]

	fmt.Println(s) // [0, 1, 2, 3, 4]

	// s1 refers to [2, 3]
	s1 := a[2:4]
	// s2 refers to [1, 2, 3, 4]
	s2 := a[1:]

	a[2] = 10

	fmt.Println(s1[0]) // 10
	fmt.Println(s2[1]) // 10

	fmt.Println(len(a))  // 5
	fmt.Println(len(s1)) // 2
	fmt.Println(cap(s1)) // 5
}
