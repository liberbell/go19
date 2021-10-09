package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}
type Group struct {
	Name   string
	People []Person
}

func main() {
	person1 := Person{"George", "Simth"}
	person2 := Person{"Elizabeth", "Plath"}
	person3 := Person{"Ronald", "Swift"}
	person4 := Person{"Lucy", "Brown"}

	people1 := []Person{person1, person2}
	people2 := []Person{person3, person4}

	gp := []Group{
		{Name: "According", People: people1},
		{Name: "Sales", People: people2},
	}

	fmt.Println(gp)
	fmt.Println(gp[0])
	fmt.Println(gp[1])
	fmt.Println(gp[1].People[0])
}
