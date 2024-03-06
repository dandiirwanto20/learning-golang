package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

//  implement struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	// struct literal
	dandi := Customer{
		Name: "Dandi",
		Address: "Jakarta",
		Age: 23,
	}

	dandi.sayHello("Hayabusa")
}
