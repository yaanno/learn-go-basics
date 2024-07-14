package main

import (
	"fmt"
)

type Product struct {
	id    int
	price float64
	title string
}

type productMap map[int]Product

func (m productMap) output() {
	fmt.Println(m)
}

func (m productMap) get(id int) Product {
	return m[id]
}

func main() {
	userNames := make([]string, 1, 5)
	userNames = append(userNames, "bob")
	userNames = append(userNames, "joe")
	fmt.Println(userNames)
	fmt.Println(len(userNames), cap(userNames))

	courses := make(map[string]string, 3)
	courses["test"] = "Test course"
	fmt.Println(courses["vue"])

	products := make(productMap)
	product := Product{1, 34, "test"}
	products[product.id] = product
	fmt.Println(products.get(1))

	for _, v := range products {
		fmt.Println("item ", v)
	}
}
