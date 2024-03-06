package main

import "fmt"

func endApp() {
	fmt.Println("End App")

	// implementasi recover benar ketika di deklarasikan di dalam func yang di defer seperti di bawah ini
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}

	// implementasi recover yang salah
	// message := recover()
	// fmt.Println("Terjadi panic", message)
}

func main() {
	// func recover biasanya digunakan untuk menangkap data dari panic
	// dengan recover proses panic akan berhenti, sehingga program akan tetap berjalan
	runApp(true)

	fmt.Println("Dandi")
}
