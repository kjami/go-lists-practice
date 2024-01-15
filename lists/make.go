package lists

import "fmt"

func printArray[T any](arr []T) {
	for i, val := range arr {
		fmt.Println("Index: ", i)
		fmt.Println("Value: ", val)
	}
}

func printMap[T comparable, K any](mapObj map[T]K) {
	for key, val := range mapObj {
		fmt.Println("Key: ", key)
		fmt.Println("Val: ", val)
	}
}

func makePractice() {
	// Give slice initial length of 3 but allocate memory for 10
	slice := make([]int, 3, 10)
	slice[0] = 1
	slice[1] = 2
	slice = append(slice, 4, 5)
	printArray(slice)

	// Give map a length 3
	productsMap := make(map[int]Product, 3)
	productsMap[1] = NewProduct(1, "Product 1", 1.99)
	productsMap[2] = NewProduct(2, "Product 2", 4.99)
	productsMap[3] = NewProduct(3, "Product 3", 5.99)

	printMap(productsMap)
}
