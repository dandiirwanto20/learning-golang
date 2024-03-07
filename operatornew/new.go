package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// go memiliki function new untuk membuat pointer
	// func new hanya mengembalikan pointer ke data kosong, artinya tidak data awal

	// implement pointer tanpa new
	var alamat1 *Address = &Address{}

	var alamat2 *Address = alamat1

	alamat1.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

	// implement new 

	alamat3 := new(Address)

	alamat4 := alamat3

	alamat4.Country = "Jimbabwe"

	fmt.Println(alamat3)
	fmt.Println(alamat4)
}
