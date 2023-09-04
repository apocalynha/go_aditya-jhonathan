# Soal Eksplorasi
## Data Dari REST API

```
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

func main() {

	productChan := make(chan []Product)

	go func() {
		resp, err := http.Get("https://fakestoreapi.com/products")
		if err != nil {
			fmt.Println("Error : ", err)
			return
		}
		defer resp.Body.Close()

		var products []Product
		err = json.NewDecoder(resp.Body).Decode(&products)
		if err != nil {
			fmt.Print("Error : ", err)
			return
		}

		productChan <- products
	}()

	products := <-productChan
	fmt.Println("Products data")
	for _, product := range products {
		fmt.Println("===")
		fmt.Printf("Title : %s\n", product.Title)
		fmt.Printf("Price : Rp.%.2f\n", product.Price)
		fmt.Printf("Category : %s\n", product.Category)
		fmt.Println("===")
	}
}
```

Output

![e_no1](/09_Concurrent-Programming/screenshots/e_no1.JPG)