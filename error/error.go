package main

import (
	"errors"
	"fmt"
)

// type error interface {
// 	Error() string
// }

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan Nol")
	} else {
		return nilai / pembagi, nil // kenapa nil karena error type datanya interface
	}
}

func main() {
	// Pada golang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interfacenya adalah error
	hasil, err := Pembagian(100, 10)
	// hasil, err := Pembagian(100, 0)

	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
