package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main() {
	// digunakan untuk menghentikan program
	// func panic dipanggil saat program terjadi panic
	// saat func panic dijalankan namun defer tetap akan di eksekusi
	runApp(true)
}
