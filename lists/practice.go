package lists

import "fmt"

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

type Product struct {
	id    int
	title string
	price float64
}

func NewProduct(id int, title string, price float64) Product {
	return Product{
		id,
		title,
		price,
	}
}

func (p Product) ToString() {
	fmt.Println(p.id, p.title, p.price)
}

func printProducts(products []Product) {
	// for i := 0; i < len(products); i++ {
	// 	products[i].ToString()
	// }
	for _, val := range products {
		val.ToString()
	}
}

func practiceExercise() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"Shuttle Badminton", "Table Tennis", "Learn new technologies"}
	fmt.Println("Hobbies Array: ", hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println("Hobbies first element: ", hobbies[0])
	fmt.Println("Hobbies 2nd and 3rd: ", hobbies[1:3])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	newHobbiesSlice1 := hobbies[:2]
	fmt.Println("New Hobbies Slice - #1: ", newHobbiesSlice1)
	newHobbiesSlice2 := hobbies[0:2]
	fmt.Println("New Hobbies Slice - #2: ", newHobbiesSlice2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	finalSlice := newHobbiesSlice2[1:3]
	fmt.Println("Final Slice: ", finalSlice)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	dynamicArr := []string{"Learn Go", "Practice everything"}
	fmt.Println("Dynamic Array: ", dynamicArr)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	dynamicArr[1] = "Takes notes and practice everything"
	dynamicArr = append(dynamicArr, "Look at official docs")
	fmt.Println("Dynamic Array - Updated: ", dynamicArr)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{
		NewProduct(1, "My First Product", 199.99),
		NewProduct(2, "My Second Product", 399.99),
	}

	fmt.Println("Products: ")
	printProducts(products)
	fmt.Println("Updated Products: ")
	products = append(products, NewProduct(3, "My Third Product", 299.99))
	printProducts(products)
}
