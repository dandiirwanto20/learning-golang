package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Nama Depan:", firstName, "nama belakang:", lastName)
}

func main()  {
	sayHello("Dandi", "Irwanto")
}
