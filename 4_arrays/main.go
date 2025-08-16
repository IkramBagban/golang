package main

import "fmt"

func main() {
	fmt.Println("-----")

	var arr1 [5]int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50

	fmt.Println("Array 1:", arr1)

	var arr2 [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array 2:", arr2)

	var arr3 = [...]int{10, 20, 30, 40, 50} // ellipsis (...) allows array length to be inferred
	fmt.Println("Array 3:", arr3)

	fmt.Println("Length of Array 3:", len(arr3))

	for i := 0; i < len(arr3); i++ {
		fmt.Printf("Element at index %d: %d\n", i, arr3[i])
	}

	fmt.Println("-----")

	// Multidimensional array
	var arr4 [2][3]int
	arr4[0][0] = 1
	arr4[0][1] = 2
	arr4[0][2] = 3
	arr4[1][0] = 4
	arr4[1][1] = 5
	arr4[1][2] = 6

	fmt.Println("Array 4:", arr4)

	for i := 0; i < len(arr4); i++ {
		for j := 0; j < len(arr4[i]); j++ {
			fmt.Printf("Element at [%d][%d]: %d\n", i, j, arr4[i][j])
		}
	}

	fmt.Println("-----")

}
