package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Dandi", "Dian", "Derry", "Dora"}
	values := []int{20, 80, 100, 7, 9, 1}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Dandi"))
	fmt.Println(slices.Index(names))
	fmt.Println(slices.Index(values))

}
