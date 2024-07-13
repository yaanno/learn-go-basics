package main

import "fmt"

type Products = []Product

type Product struct {
	title string
	id    int
	price float64
}

func newProduct(title string, id int, price float64) Product {
	return Product{
		title: title,
		id:    id,
		price: price,
	}
}

func (p Product) formatted() string {
	return fmt.Sprintf(
		"Title: %v\nPrice: $%v",
		p.title,
		p.price,
	)
}

func main() {
	// 1
	hobbies := [3]string{
		"golf",
		"programming",
		"fencing",
	}
	fmt.Println(hobbies)
	// 2
	fmt.Println(hobbies[0])

	lasttwo := [2]string{hobbies[1], hobbies[2]}
	fmt.Println(lasttwo)
	// 3
	slice1 := hobbies[:len(hobbies)-1]
	slice2 := hobbies[:2]
	fmt.Println(slice1, slice2)

	// 4
	reslice := slice1[1:3]
	fmt.Println(reslice)

	// 5
	dynamic := []string{}
	dynamic = append(
		dynamic,
		"become a pro go developer",
		"become a go senior programmer",
	)
	fmt.Println(dynamic)

	// 6
	dynamic[1] = "become a javascript senior dev"
	fmt.Println(dynamic)
	dynamic = append(dynamic, "Learn more!")
	fmt.Println(dynamic)

	// bonus
	products := Products{
		Product{
			title: "Go advanced book",
			id:    1,
			price: 65},
		Product{
			title: "Go services in practice",
			id:    2,
			price: 23},
	}
	fmt.Println(products)
	products = append(products, Product{
		title: "Testing in go",
		id:    3,
		price: 67})
	fmt.Println(products)

	productsList := Products{
		newProduct("Test", 5, 34.0),
		newProduct("Test2", 5, 34.0),
	}
	fmt.Println(productsList)
	productsList = append(
		productsList,
		newProduct("Test3", 5, 34.0),
	)
	fmt.Println(productsList)
	fmt.Println(productsList[0].formatted())
}
