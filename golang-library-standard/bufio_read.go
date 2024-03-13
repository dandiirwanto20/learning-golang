package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	// Reader (Input)
	input := strings.NewReader("This is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()

		// seleksi ketika end of file
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
