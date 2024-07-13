package main

import "fmt"

type Product struct {
	id    int
	price float64
	title string
}

func main() {
	websites := map[string]string{
		"aws":    "http://aws.com",
		"google": "https://google.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["aws"])
	websites["linkedin"] = "https://linkedin.com"
	fmt.Println(websites)
	delete(websites, "google")
	fmt.Println(websites)
	product := Product{1, 23.9, "my fav book"}
	var products = map[int]Product{}
	products[product.id] = product
	fmt.Println(products[1])
}
