package main

import (
	"fmt"
)

type validationError struct {
	Message string
}

type notFoundError struct {
	MessErr string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.MessErr
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "dandi" {
		return &notFoundError{MessErr: "Data Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("dandi", nil)
	if err != nil {
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown error", err.Error())
		// }
		
		// dengan switch case
		switch finalError := err.(type) {
			case *validationError:
				fmt.Println("validation error:", finalError.Error())
			case *notFoundError:
				fmt.Println("not found error:", finalError.Error())
			default:
				fmt.Println("unknown error")
		}
	} else {
		fmt.Println("Sukses")
	}
}
