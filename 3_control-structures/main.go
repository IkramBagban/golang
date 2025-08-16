package main

import "fmt"

func main() {
	fmt.Println("-----")

	var age int = 20

	if age >= 18 {
		fmt.Println("Adult")
	} else {

		fmt.Println("Minor")
	}

	fmt.Println("-----")

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----")
	var day int = 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednessday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	fmt.Println("-----")
}
