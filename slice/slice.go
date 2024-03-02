package main

import "fmt"

func main() {
	names := [...]string{
		"Dandi",
		"Irwanto",
		"Rusman",
		"Hena",
		"Joko",
		"Tio",
	}

	slice := names[4:6] // pointer 4, length 2 (antara 4-6), capacity 2
	fmt.Println(slice)

	slice2 := names[:3] //0,1,2
	fmt.Println(slice2)

	slice3 := names[3:] //3,4,5
	fmt.Println(slice3)

	slice4 := names[:] //all
	fmt.Println(slice4)

	// manual data type slice
	// var slice4 []string = names[:]

	// slice function
	// len panjang dari slice
	fmt.Println(len(slice))

	// capacity dari slice
	fmt.Println(cap(slice))

	// append slice -> membuat slice baru dengan memasukan data ke posisi terakhir slicenya, jika kapasistas sudah penuh, maka akan membuat array baru
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daySlice1 := days[5:] // Sabtu Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"  // Mengubah sabtu menjadi sabtu baru
	daySlice1[1] = "Minggu Baru" // mengubah minggu menjadi minggu baru

	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru") // append slice
	daySlice2[0] = "0 Ups" // Mengubah data sabtu baru menajadi 0 ups 

	fmt.Println(daySlice2) // mengambil data slice dari dataSlice1
	fmt.Println(days) // kenapa tidak berubah ketika di append karena capasitas array sudah penuh
}
