package main

import "fmt"

func main() {
	/*
		sebuah anonymous function atau function tanpa nama
		yang dapat disimpan sebagai sebuah variable maupun
		dapat dijadikan sebagai nilai return dari sebuah function
	*/

	// perlu diperhatikan jangan terlalu sering menggunakan closure ketika program sudah kompleks karena akan membingungkan nantinya
	counter := 0

	// bayangkan code sudah banyak

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)

}
