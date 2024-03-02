package main

import "fmt"

func main() {
	// slice dari awal array otomatis dibuat dari slice
	newSlice := make([]string, 2, 5) // panjang 2, 5 kapasitas

	newSlice[0] = "Dandi"
	newSlice[1] = "Irwanto"
	// newSlice[2] = "Irwanto" akan error kalo melebihi panjang "panic: runtime error: index out of range [2] with length 2"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) // check length
	fmt.Println(cap(newSlice)) // check capacity

	// solution:
	newSlice2 := append(newSlice, "Dand")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2)) // check length
	fmt.Println(cap(newSlice2)) // check capacity

	newSlice2[0] = "Ir"

	fmt.Println(newSlice2)
	fmt.Println(newSlice) // condition [Ir Irwanto] karena masih menggunakan array yang sama karena kapasitas masih bisa menampung, beda kondisi ketika kapasitas penuh maka akan membuat array baru
}
