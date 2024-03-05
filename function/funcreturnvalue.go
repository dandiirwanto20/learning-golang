package main

import "fmt"

func getHello(name string) string { // param name dengan tipe data string dan pengembalian data adalah string
	return "Hello " + name
}

func sumNumber(num int) int {
	return num + num
}
func main() {
	result := getHello("Dandi")
	fmt.Println(result)

	// sum data
	fmt.Println(sumNumber(4))
}