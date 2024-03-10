package main

import (
	"fmt"
	"os"
)

func main() {
	// contoh Args
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}
	// package Hostname
	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

}
