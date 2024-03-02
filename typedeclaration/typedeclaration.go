package main

import "fmt"

func main() {
	// type declaration type data baru
	type NoKTP string

	var ktpDan NoKTP = "111111111"

	var contoh string = "22222222"
	var contohKTP NoKTP = NoKTP(contoh)
	fmt.Println(ktpDan)
	fmt.Println(contohKTP)
}
