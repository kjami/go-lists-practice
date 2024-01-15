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
}
