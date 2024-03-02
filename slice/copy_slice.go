package main

import "fmt"

func main() {
	// array days
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// different about array and slice
	iniArray := [...]int{
		1,
		2,
		3,
		4,
		5,
	}

	iniSlice := []int{
		1,
		2,
		3,
		4,
		5,
	}

	fmt.Println(iniArray)
	// fmt.Printf("Type of Variable is : %T\n", iniArray) //check type
	fmt.Println(iniSlice)
	// fmt.Printf("Type of Variable is : %T\n", iniSlice) //check type
}
