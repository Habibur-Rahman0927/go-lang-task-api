package array

import "fmt"

func Array() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var arr3 = [...]int{1, 2, 3}
	arr4 := []int{4, 5, 6, 7, 8}

	fmt.Println(arr3)
	fmt.Println(arr4)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)
}
