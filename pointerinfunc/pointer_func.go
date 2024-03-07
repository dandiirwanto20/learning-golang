package main

import "fmt"

type Address struct {
	City, Province, Country string
}
// with pointer in param
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

// without pointer
// func ChangeCountryToIndonesia(address Address) {
// 	address.Country = "Indonesia"
// }

func main() {
	// pointer di function digunakan saat ingin mengubah data parameter asli di function (pass by reference)

	// case
	address := &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	ChangeCountryToIndonesia(address)

	fmt.Println(address) // data kosong tidak expected masih pass by reference
}
