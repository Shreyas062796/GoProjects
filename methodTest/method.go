package main

import (
	"fmt"
)

func main() {
	personObject := Person{
		Name: "Shreyas",
		Age: 23,
		Occupation: "Software Engineer",
	}
	personObject.greet()
	changeAge(&personObject,10)
	personObject.greet()
}

func (p Person) greet() {
	fmt.Println(p.Name, ":", p.Age,":",p.Occupation)
}

func changeAge(p *Person, newAge int) {
	p.Age = newAge
}
type Person struct {
	Name string
	Age int
	Occupation string
}