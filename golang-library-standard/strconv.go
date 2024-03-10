package main

import (
	"fmt"
	"strconv"
)

func main() {
	// untuk melakukan conversi type data

	// contoh convert bool
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	// contoh convert int
	resultInt, err := strconv.Atoi("1000")
	if err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("Error", err.Error())
	}

	// convert binary
	binary := strconv.FormatInt(999, 2)

	fmt.Println(binary)

	// convert int to string
	var stringInt string = strconv.Itoa(999)

	fmt.Println(stringInt)

}
