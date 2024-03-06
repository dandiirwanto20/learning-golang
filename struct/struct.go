package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	// template data yang digunakan untuk menggabungkan nol atau lebih type data dalam satu kesatuan
	// Struct adalah kumpulan dari field
	// Data dalam struct disimpan dalam field

	// implementasi struct as a type

	var dan Customer
	dan.Name = "Dandi Irwanto"
	dan.Address = "Rembang"
	dan.Age = 20

	fmt.Println(dan)
	fmt.Println(dan.Name)
	fmt.Println(dan.Address)
	fmt.Println(dan.Age)

	// implement struct literal
	dandi := Customer{
		Name: "Dan",
		Address: "Jakarta",
		Age: 23,
	}

	fmt.Println(dandi)

	// OR
	ir := Customer{"Dan", "Cibinong", 24}

	fmt.Println(ir)
}
