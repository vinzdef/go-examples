package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Sayan struct {
	Person
	Power int
}

func (p Person) greet() {
	fmt.Println("I'm", p.Name)
}

func (p *Person) say(word string) {
	fmt.Println(p.Name, "says:", word)
}

func (s *Sayan) damage() {
	fmt.Println(s.Name, "made", s.Power, "damage")
}

func main() {
	gohan := Sayan{Person{"Gohan", 12}, 9000}
	gohan.greet()
	gohan.say("I'm my father's son!")
	gohan.damage()
}
