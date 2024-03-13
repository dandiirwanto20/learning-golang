package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` // structTag
}

type Person struct {
	Name, Address, Email string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)

	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	t := reflect.TypeOf(value)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return true
}

func main() {
	readField(Sample{"Dandi"})
	readField(Person{"Dandi", "", ""})

	person := Person{
		Name:    "Dandi",
		Address: "Rembang",
		Email:   "dandi.ir20@gmail.com",
	}

	fmt.Println(IsValid(person))
}
