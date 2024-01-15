package lists

import "fmt"

func ExecuteLists() {
	var arr [4]int = [4]int{1, 2, 3, 4}
	strArr := [...]string{"one", "two", "three", "four", "five", "six"}
	boolArr := [2]bool{true, false}
	fmt.Println(arr, strArr, boolArr)
	arr[1] = -2
	fmt.Println(arr)
	strSubArr := strArr[1:]
	fmt.Println(strSubArr)
	strSubArr2 := strSubArr[:4]
	fmt.Println(strSubArr2)
	lastStrArr := strSubArr2[1:3]
	fmt.Println(lastStrArr)

	// To understand the behaviour of len and cap on arrays and slices
	// Case 1: slicing from left
	var intArr [5]int = [5]int{1, 2, 3, 4, 5}
	firstSlicedArr := intArr[1:] //2, 3, 4, 5
	fmt.Println(len(firstSlicedArr), cap(firstSlicedArr), firstSlicedArr)
	secondSlicedArr := firstSlicedArr[1:2] // 3
	fmt.Println(len(secondSlicedArr), cap(secondSlicedArr), secondSlicedArr)
	// Perform slice on the same slice but provide full capacity on the right (and empty value on the left)
	secondSlicedArr = secondSlicedArr[:3] // 3, 4, 5
	fmt.Println(len(secondSlicedArr), cap(secondSlicedArr), secondSlicedArr)

	// Case 2: slicing to the right
	var anotherStrArr [5]string = [5]string{"one", "two", "three", "four", "five"}
	firstSlicedStrArr := anotherStrArr[:3] //"one", "two", "three"
	fmt.Println(len(firstSlicedStrArr), cap(firstSlicedStrArr), firstSlicedStrArr)
	secondSlicedStrArr := firstSlicedStrArr[1:2] // "two"
	fmt.Println(len(secondSlicedStrArr), cap(secondSlicedStrArr), secondSlicedStrArr)
	// Perform slice on the same slice but provide full capacity on the right (and empty value on the left)
	secondSlicedStrArr = secondSlicedStrArr[:4] // "two", "three", "four", "five"
	fmt.Println(len(secondSlicedStrArr), cap(secondSlicedStrArr), secondSlicedStrArr)
	firstSlicedStrArr = firstSlicedStrArr[:5] // "one", "two", "three", "four", "five"
	fmt.Println(len(firstSlicedStrArr), cap(firstSlicedStrArr), firstSlicedStrArr)

	// Initialize array without specifying length
	var dynamicArr []int = []int{1, 2, 3}
	// Append additional items - This will create a new array
	dynamicArr = append(dynamicArr, 4, 5, 6)
	fmt.Println(dynamicArr) // 1, 2, 3, 4, 5, 6
}
