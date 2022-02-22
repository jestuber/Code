package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"` // this is a tag
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{}) // need reflect library to get tag
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
