package main

import "fmt"

type Man struct {
	Name string
}

// implement without pointer
// func (man Man) Maried() { // param with value not preference
// 	man.Name = "Mr. " + man.Name
// }

func (man *Man) Maried() { // param with value not preference
	man.Name = "Mr. " + man.Name
}

func main() {
	// method menempel pada struct, tapi pada dasarnya pass by value (di copy datanya dan diduplikat)
	//sangat direkomendasikan menggunakan pointer pada method agar tidak boros dalam penggunaan memory

	// case method tanpa pointer
	dandi := Man{"Dandi"} // implement struct literal

	dandi.Maried()

	fmt.Println(dandi.Name)

}
