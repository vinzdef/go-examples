package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

var Chicco = Person{"Chicco", 28}
var Princess = Person{
	Name: "Princess",
	Age:  23,
}

type Sayan struct {
	Person // same as Person Person
	Power  int
}

var Gohan = Sayan{
	Person{"Gohan", 13},
	9001,
}

fmt.Println(Gohan.Name) // "Gohan"
// Same as Gohan.Person.Name

func main() {
	fmt.Println(Chicco)
	fmt.Println(Gohan)
}
