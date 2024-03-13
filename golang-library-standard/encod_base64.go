package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var encoded = base64.StdEncoding.EncodeToString([]byte("Dandi Irwanto"))

	// proses encoding string to base64
	fmt.Println(encoded)
	var decoded, err = base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
