package main

import (
	"fmt"
	"strings"
)

func main() {
	//package untuk manipulasi data string
	fmt.Println(strings.Contains("Dandi Irwanto", "Dandi"))              //mengecek apakah mengandung string lain
	fmt.Println(strings.Split("Dandi Irwanto", " "))                     // return value is slice (memotong string berdasarkan separator)
	fmt.Println(strings.ToLower("DANDI IRWANTO"))                        // mengubah huruf menjadi kecil
	fmt.Println(strings.ToUpper("dandi irwanto"))                        // mengubah huruf menjadi besar atau kapital
	fmt.Println(strings.Trim("   Dandi   ", " "))                        // menghapus/cutset awal dan akhir string
	fmt.Println(strings.ReplaceAll("Irwanto Dandi", "Irwanto", "Dandi")) // mengubah semua string dari from ke to
}
