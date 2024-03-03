package main

import "fmt"

func main() {
	name := "Dandi"

	switch name {
	case "Dandi":
		fmt.Println("Hallo Dandi")
	case "Budi":
		fmt.Println("Hay Budi")
	default:
		fmt.Println("Salah")
	}

	// switch dengan short statement
	switch length := len(name); length > 6 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama Benar")
	}

	// switch tanpa kondisi tapi agar lebih familiar disarankan untuk memakai ifelse saja
	length1 := len(name)

	switch {
		case length1 > 6:
			fmt.Println("Nama terlalu panjang oy")
		case length1 == 5:
			fmt.Println("Nama adalah benar")
		default:
			fmt.Println("Nama default")
	}
}
