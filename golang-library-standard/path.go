package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {

	// implement path
	pathValues := "hello/world.go"
	fmt.Println(path.Dir(pathValues))
	fmt.Println(path.Base(pathValues))
	fmt.Println(path.Ext(pathValues))
	fmt.Println(path.Join("hello", "world", "main.go"))

	// implement filepath
	fmt.Println(filepath.Dir(pathValues))
	fmt.Println(filepath.Base(pathValues))
	fmt.Println(filepath.Ext(pathValues))
	fmt.Println(filepath.IsAbs(pathValues))
	fmt.Println(filepath.IsLocal(pathValues))
	fmt.Println(filepath.Join("helllo", "world", "main.go"))
}
