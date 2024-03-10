package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))  // membulatkan float64 ke atas
	fmt.Println(math.Floor(1.60)) // membulatkan float64 ke bawah
	fmt.Println(math.Round(1.60)) // membulatkan float64 ke atas atau ke bawah yang paling dekat
	fmt.Println(math.Max(10, 11)) // mengembalikan nilai yang paling besar
	fmt.Println(math.Min(10, 11)) // mengembalikan nilai yang paling kecil
}
