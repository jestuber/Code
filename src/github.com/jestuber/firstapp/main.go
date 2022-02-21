package main

import (
	"fmt"
)

func main() {
	//grades := [3]int{97, 85, 93}
	grades := [...]int{97, 85, 93} // ... gets length from data
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students))

	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := []int{1, 2, 3} // slice
	d := c              // slices are reference types, this does not make a copy
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("Length: %v\n", len(c))
	fmt.Printf("Capacity: %v\n", cap(d))

	t := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	u := t[:]   // slice of all elements
	v := t[3:]  // slice from 4th ele to end
	x := t[:6]  // slice first 6 ele
	y := t[3:6] // slice 4,5,and 6th ele
	fmt.Println(t)
	fmt.Println(u)
	fmt.Println(v)
	fmt.Println(x)
	fmt.Println(y)

	w := make([]int, 3, 100)
	fmt.Println(w)
	fmt.Printf("Length: %v\n", len(w))
	fmt.Printf("Capacity: %v\n", cap(w))

	z := []int{}
	fmt.Println(z)
	z = append(z, 1)
	fmt.Println(z)
	//z = append(z, 2, 3, 4, 5)
	z = append(z, []int{2, 3, 4, 5}...)
	fmt.Println(z)
}
