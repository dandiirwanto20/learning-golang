package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Dandi"
	middleName = "Ir"
	lastName = "Irwanto"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()

	fmt.Println(firstName, middleName, lastName)
}
