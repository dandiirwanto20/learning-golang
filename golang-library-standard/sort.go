package main

import (
	"fmt"
	"sort"
)

// membuat struct
type User struct {
	Name string
	Age  int
}

// slice
type UserSlice []User

// membuat kontrak untuk len, less, dan swap
func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	// temp := s[i]
	// s[i] = s[j]
	// s[j] = temp

	// di golang ada cara yang lebih gampang untuk swap data
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Dand", 25},
		{"Dian", 20},
		{"Dion", 35},
		{"Deo", 15},
	}

	sort.Sort(UserSlice(users)) // implement sort.Sort untuk mengurutkan data

	fmt.Println(users)
}
