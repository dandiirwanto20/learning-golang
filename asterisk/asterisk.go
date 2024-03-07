package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Operator *
	address1 := Address{"Cibinong", "Jakarta", "Indonesia"}

	address2 := &address1 // copy value address1 bukan me replace

	address2.City = "Bandung"

	fmt.Println(address1) // data ikut berubah
	fmt.Println(address2) // data berubah

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address1)
	// fmt.Println(address2)

	// implement operator * untuk membawa value yang diacu sebelumnya
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // implement operator *
	fmt.Println(address1)
	fmt.Println(address2)

}
