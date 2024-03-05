package main

import "fmt"

// type declaration
type Filter func(string) string

func sayHellSpam(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHellSpam("Dandi", spamFilter)

	spam := spamFilter

	sayHellSpam("Anjing", spam)
}
