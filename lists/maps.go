package lists

import "fmt"

func maps() {
	// init a map with key of int type and value of Product type
	productsMap := map[int]Product{
		1: NewProduct(1, "Product 1", 1.99),
		2: NewProduct(2, "Product 2", 2.99),
	}

	fmt.Println("Products Map: ", productsMap)
	// Add new elements - new map (and new memory) will be allocated
	productsMap[3] = NewProduct(3, "Product 3", 4.99)
	fmt.Println("Products Map - Updated: ", productsMap)
	// delete one of the map items
	delete(productsMap, 2)
	fmt.Println("Products Map - After deletion: ", productsMap)
}
