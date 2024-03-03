package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Println("Selesai")

	// For dengan statement, 2 statement (init dan post statement)
	for counter1 := 1; counter1 <= 10; counter1++ {
		fmt.Println("Perulangan", counter1)
	}

	// for Range, iterasi untuk data collection seperti Array, Slice, Map
	names := []string{
		"Dandi",
		"Irwanto",
		"Dean",
		"Ruli",
	}
	// for manual
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range untuk contoh slice
	for index, name := range names {
		fmt.Println("index", index, "name", name)
	}

	// for range untuk contoh array
	address := [...]string{
		"Rembang",
		"Jakarta",
		"Palembang",
		"Jogja",
	}

	for _, name1 := range address { // jika tidak butuh index bisa menggunakan _
		fmt.Println("Alamat", name1)
	}

	// for range untuk map

	person := map[string]string{
		"name": "Dandi",
		"address":"Jakarta",
		"age":"23",
	}

	for key, value := range person {
		fmt.Println("key", key, "value", value)
	}
}
