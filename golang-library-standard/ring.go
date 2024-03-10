package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// circular list
	data := ring.New(5)

	// make manually
	// data.Value = "Value 1"

	// data = data.Next()
	// data.Value = "Value 2"
	// data = data.Next()
	// data.Value = "Value 3"
	// data = data.Next()
	// data.Value = "Value 4"
	// data = data.Next()
	// data.Value = "Value 5"

	// make with loop
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
