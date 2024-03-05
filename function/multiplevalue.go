package main

import "fmt"

func getFullName() (string, string) {
	return "Dandi", "Irwanto"
}

func main() {
	firstName, lastName := getFullName()

	fmt.Println(firstName, lastName)

	// menghiraukan return value ketika multiple
	namaDepan, _ := getFullName()

	fmt.Println(namaDepan)
}
