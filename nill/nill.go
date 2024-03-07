package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// dalam golang golang variable yang belum di inisialisasi maka akan dibuatkan default valuenya
	// tetapi dalam go juga ada nill yaitu data kosong

	// support untuk tipe data interface, function, map, slice, pointer, dan channel
	data := NewMap("Dandi")

	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
