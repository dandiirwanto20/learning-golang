package main

import "fmt"

func main() {
	const firstName string = "Dandi" 
	const lastName = "Irwanto"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// multiple const

	const (
		fName = "Dand"
		lName = "Ir"
	)

	fmt.Println(fName)
	fmt.Println(lName)
}
