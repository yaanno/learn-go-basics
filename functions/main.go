package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	doubler := createTransformer(2)
	tripler := createTransformer(3)
	transformed := transformNumbers(
		&numbers,
		func(number int) int { return number * 2 },
	)
	fmt.Println(transformed)
	doubled := transformNumbers(&numbers, doubler)
	tripled := transformNumbers(&numbers, tripler)
	fmt.Println(doubled, tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}
