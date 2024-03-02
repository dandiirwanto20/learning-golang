package main

import "fmt"

func main() {
	// map menggunakan key-value
	person := map[string]string{
		"name":    "Dandi",
		"address": "Jakarta",
	}

	// manual declaration
	// 	var persons map[string]string = map[string]string
	// 	persons["name"] = "Dandi"
	// 	persons["address"] = "Rembang"

	fmt.Println(person["name"])    // ambil data name
	fmt.Println(person["address"]) // ambil data address
	fmt.Println(person)
	// fmt.Println(person["invalid"]) error string kosong karena tidak ada key

	//map function

	// len map
	fmt.Println(len(person))

	// ubah data map
	person["address"] = "Bali"
	fmt.Println(person["address"])

	//membuat map baru menggunakan make
	// book := map[string]string{}
	book := make(map[string]string)

	book["title"] = "Buku Sigma"
	book["author"] = "Dand"
	book["wrong"] = "Ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
