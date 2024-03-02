package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "D"
	names[1] = "a"
	names[2] = "n"

	fmt.Print(names[0])
	fmt.Print(names[1])
	fmt.Print(names[2])
	fmt.Println("")
	fmt.Print(names)
	fmt.Println("")
	fmt.Println("===========")

	// array langsung
	values := [3]int{
		40,
		50,
		60,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}
