package main

import "fmt"

func getGoodbye(name string) string {
	return "Good bye " + name
}

func main() {
	goodbye := getGoodbye // memanggil function as value

	fmt.Println(goodbye("Dandi"))
}
