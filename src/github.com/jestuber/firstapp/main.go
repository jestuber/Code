package main

import "fmt"

func main() {
	aDoctor := struct{ name string }{name: "John Pertwee"} // anonymous struct
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker" // structs are not reference types, this will not modify aDoctor
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

	yetAnotherDoctor := &aDoctor
	yetAnotherDoctor.name = "Joe Smith"
	fmt.Println(aDoctor)
	fmt.Println(yetAnotherDoctor)
}
