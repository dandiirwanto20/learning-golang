package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put yaour database password")

	flag.Parse()

	fmt.Println(*host, *username, *password) //barikan dari flag adalah pointer
}
