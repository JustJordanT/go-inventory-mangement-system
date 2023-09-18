package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type ProductCollection map[int]Product

func main() {
	inventory := make(ProductCollection)

	inventory.addItem(Product{ID: 1, Name: "Product 1", Price: 1.0, Quantity: 6})

	//updateProduct(Product{ID: 1, Name: "Product 1", Price: 10.0, Quantity: 11}, mapProduct)

	jsonData, err := fileReader("inventory.json")
	if err != nil {
		fmt.Println(err)
	}

	err = importInventoryFromJson(jsonData, inventory)
	if err != nil {
		fmt.Println(err)
	}

	inventory.viewProduct(inventory)

	fmt.Println("SEARCHING BY ID..")
	inventory.viewProductByPriceRange(1.0, 10.0)

	fmt.Println("REPORTING..")
	fmt.Println(inventory.totalInventoryValue())
	fmt.Println(inventory.lowQuantityProducts())

}

func (pc ProductCollection) addItem(product Product) {
	pc[product.ID] = product
}

func (pc ProductCollection) updateProduct(product Product) {
	pc[product.ID] = product
}

func (pc ProductCollection) deleteProduct(id int) {
	delete(pc, id)
}

func (pc ProductCollection) viewProduct(collection ProductCollection) {
	for _, product := range collection {
		fmt.Printf("ID: %v\n", product.ID)
		fmt.Printf("Name: %v\n", product.Name)
		fmt.Printf("Price: %v\n", product.Price)
		fmt.Printf("Quantity: %v\n", product.Quantity)
		fmt.Println()
	}
}

func (pc ProductCollection) viewProductByID(id int) {
	fmt.Println(pc[id])
}

func (pc ProductCollection) viewProductByName(name string) {
	for _, product := range pc {
		if product.Name == name {
			fmt.Println(product)
		}
	}
}

func (pc ProductCollection) viewProductByPrice(price float64) {
	for _, product := range pc {
		if product.Price == price {
			fmt.Println(product)
		}
	}
}

func (pc ProductCollection) viewProductByQuantity(quantity int) {
	for _, product := range pc {
		if product.Quantity > quantity {
			fmt.Println(product)
		}
	}
}

func (pc ProductCollection) viewProductByPriceRange(start float64, end float64) {
	for _, product := range pc {
		if product.Price >= start && product.Price <= end {
			fmt.Println(product)
		}
	}
}

// Reporting Functions
func (pc ProductCollection) totalInventoryValue() float64 {
	var total float64
	for _, product := range pc {
		total += product.Price * float64(product.Quantity)
	}

	return total
}

func (pc ProductCollection) mostExpensiveProduct() Product {
	var mostExpensive Product
	for _, product := range pc {
		if product.Price > mostExpensive.Price {
			mostExpensive = product
		}
	}

	return mostExpensive
}

func (pc ProductCollection) leastExpensiveProduct() Product {
	var leastExpensive Product
	for _, product := range pc {
		if product.Price < leastExpensive.Price {
			leastExpensive = product
		}
	}

	return leastExpensive
}

func (pc ProductCollection) lowQuantityProducts() Product {
	var lowestQuantity Product
	for _, product := range pc {
		if product.Quantity < 100 {
			lowestQuantity = product
		}
	}

	return lowestQuantity
}

// Export/Import Functions

func (pc ProductCollection) exportInventoryToJson() (string, error) {
	jsonData, err := json.MarshalIndent(pc, "", "    ")
	if err != nil {
    return "", fmt.Errorf("error: %v", err)
	}

	// Print the JSON data
	return string(jsonData), nil

}

func importInventoryFromJson(jsonData []byte, inventory ProductCollection) error {
	// TODO Import Inventory from JSON File Need to create a file reader function to read the json file as a byte array
	var objects []Product

	// Unmarshal the JSON data into the slice of Product structs
	if err := json.Unmarshal(jsonData, &objects); err != nil {
		return err
	}

	// Loop through the slice of products and add each one to the inventory map
	for _, product := range objects {
		inventory[product.ID] = product
	}

	return nil
}

func exportInventoryToCsv(mapProduct map[int]Product) {
	// TODO Export Inventory to CSV
}

func importInventoryFromCsv(mapProduct map[int]Product) {
	// TODO Import Inventory from CSV
}

func fileReader(file string) ([]byte, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
