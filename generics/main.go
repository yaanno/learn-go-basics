package main

import "fmt"

func main() {
	sum := add(1.6, 5.7)
	fmt.Println(sum)
}

func add[T int | string | float64](a, b T) T {
	return a + b
}
