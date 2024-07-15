package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	otherNumbers := []int{5, 2, 3, 4}
	transformerFn := getTransformerFn(&numbers)
	otherTransformerFn := getTransformerFn(&otherNumbers)
	doubled := transformNumbers(&numbers, transformerFn)
	tripled := transformNumbers(&numbers, otherTransformerFn)

	fmt.Println(doubled, tripled)
}
func transformNumbers(numbers *[]int, f transformFn) []int {
	dNumbers := []int{}
	for _, v := range *numbers {
		dNumbers = append(dNumbers, f(v))
	}
	return dNumbers
}

func getTransformerFn(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
