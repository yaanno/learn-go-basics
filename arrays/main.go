package main

import "fmt"

func main() {
	prices := []float64{12, 13}
	otherslice := []float64{234, 234, 234}
	prices = append(prices, otherslice...)
	fmt.Println(prices)
	fmt.Println(prices[0])
}
