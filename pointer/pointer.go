package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// implement pass by value
	address1 := Address{"Cibinong", "Jakarta", "Indonesia"}

	address2 := address1 // copy value address1 bukan me replace

	address2.City = "Bandung"

	fmt.Println(address1) // data tidak berubah
	fmt.Println(address2) // data berubah

	// implement pointer untuk pass by reference untuk replace data memory untuk menghindari penggunaan memory yang berlebihan

	// operator & diikuti nama variabel
	address3 := Address{"Jambi", "Lampung", "Indonesia"}

	// implement var manually
	// var address3 Address = Address{"Jambi", "Lampung", "Indonesia}
	// var address4 *Address = &address3 

	address4 := &address3 // implement pointer

	address4.City = "Lumajang"

	fmt.Println(address3)
	fmt.Println(address4)

}
