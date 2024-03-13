package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout) // tujuan dari output
	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("selamat belajar\n")

	writer.Flush()
}
