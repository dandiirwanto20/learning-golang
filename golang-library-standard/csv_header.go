package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	// csv encoding
	csvString := "dandi,irwanto\n" +
		"dian,sastro\n" +
		"joko,jian,jordi"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
