package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// anonymous func dalam sebuah var
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Dandi", blacklist)
	// anonymous func di mana case sederhana
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
