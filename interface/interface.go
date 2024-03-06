package main

import "fmt"

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

type HasName interface {
	GetName() string // kontrak
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func (person Person) GetName() string { // kontrak interface
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	// interface type data abstract, tidak memiliki implementasi langsung
	// berisikan sebuah definisi-definisi method
	// interface digunakan sebagai kontrak

	// implement interface
	person := Person{Name: "Dandi"}
	SayHello(person)

	animal :=  Animal{Name: "Kucing"}
	SayHello(animal)
}
