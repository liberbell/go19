package main

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
}
