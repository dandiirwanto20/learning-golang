package main

import "fmt"

func main() {
	// menentukan array tidak pasti bisa menggunakan `...`
	values := [...]int{
		10,
		20,
		30,
		40,
		50,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	
	// mengubah data array
	values[0] = 11
	fmt.Println(values)
}
