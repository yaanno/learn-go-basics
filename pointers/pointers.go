package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("Age:", *agePointer)
	editAgetToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgetToAdultYears(age *int) {
	*age = *age - 18
}
