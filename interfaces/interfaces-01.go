package main

import "fmt"

type Greetable interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return fmt.Sprint("Person ", p.Name)
}

type Animal struct {
	Nickname string
}

func (a Animal) Greet() string {
	return fmt.Sprint("Animal ", a.Nickname)
}

func main() {
	p := Person{"Chicco"}
	a := Animal{"Wolfie"}

	meet(p, a)
}

func meet(g1 Greetable, g2 Greetable) {
	fmt.Println(g1.Greet(), "meets", g2.Greet())
}
