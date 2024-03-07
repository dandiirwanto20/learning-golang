package main

import "fmt"

func Ups() interface{} {
	// return 1
	// return true

	return "Ups"
}

// func menggunakan return value any pada versi golang terbaru
func Wupsy() any {
	return "Upsy"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)

	anyKosong := Wupsy()
	fmt.Println(anyKosong)
}
