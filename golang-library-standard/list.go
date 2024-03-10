package main

import (
	"container/list"
	"fmt"
)

func main() {
	// case: membuat structured data linked list (biasanya untuk antrian)
	var data *list.List = list.New()

	data.PushBack("Dandi")
	data.PushBack("Irwanto")
	data.PushBack("Ir")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	// if make with loop
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
