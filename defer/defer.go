package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")

	// implemetasi jika tidak menggunakan defer maka dia sync
	// jika code sebelumnya error maka func logging tidak akan dieksekusi
	// logging()
}

func main() {
	// merupakan sebuah function yang bisa kita jadwalkan untuk dieksekusi setelah function selesai di eksekusi
	// seperti implementasi try catch secara async
	runApplication()

}
