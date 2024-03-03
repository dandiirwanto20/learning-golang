package main

import "fmt"

func main() {
	name := "Dandi"

	// if expression
	if name == "Dandi" {
		fmt.Println("Hello", name)
	} else { // else expression
		fmt.Println("kondisi salah")
	}

	// else if expression
	address := "Rembang"

	if address == "Rembang" {
		fmt.Println("Rumah", address)
	} else if address == "Jakarta" {
		fmt.Println("Alamat Salah")
	} else {
		fmt.Println("Kondisi Salah")
	}

	// if short statement
	if length := len(name); length > 6 {
		fmt.Println("Tidak lebih dari enam")
	} else {
		fmt.Println("Nama Benar", len(name))
	}
}
