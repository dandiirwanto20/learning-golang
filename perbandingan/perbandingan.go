package main

import "fmt"

func main() {
	name1 := "Dan"
	name2 := "Dan"

	// ==
	result := name1 == name2

	fmt.Println(result)

	// !=
	result = name1 != name2
	fmt.Println(result)

	// >
	name1 = "b"
	name2 = "a"

	result = name1 > name2
	fmt.Println(result)

	// <
	result = name1 < name2
	fmt.Println(result)

	// >=
	result = name1 >= name2
	fmt.Println(result)

	// <=
	result = name1 <= name2
	fmt.Println(result)
}
