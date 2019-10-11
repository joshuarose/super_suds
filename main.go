package main

import (
	"fmt"
)

// Product is a representation of products P&G sells
type Product struct {
	Name string
}

func main() {
	products := [3]Product{
		Product{Name: "Warsh Cloth"},
		Product{Name: "Soap Bar"},
		Product{Name: "Soap Bucket"},
	}

	for _, product := range products {
		fmt.Println(product.Name)
	}
}
