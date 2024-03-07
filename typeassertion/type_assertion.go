package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	// type assertion adalah kemampuan mengubah tipe data yang kita inginkan
	// sering digunakan ketika menggunakan interface kosong (interface{} / any) bisa dibilang conversion any ke tipe data yang diinginkan
	result := random()

	resultString := result.(string)

	fmt.Println(resultString)

	// hati-hati ketika convert, harus dipastikan type datanya benar atau tidak karena kalau salah akan terjadi panic
	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)

	// cara lebih aman untuk menggunakan type assertion dengan switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("Unknown")
	}
}
