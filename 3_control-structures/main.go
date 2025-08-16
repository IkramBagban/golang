package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-----")

	var age int = 20

	if age >= 18 {
		fmt.Println("Adult")
	} else {

		fmt.Println("Minor")
	}

	fmt.Println("----- loop")

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("----- switch")
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

	fmt.Println("----- multiple condition switch")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // we can add multiple conditoin like this.
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's workday")
	}

	fmt.Println("----- type switch")
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I'm an int:", t)
		case string:
			fmt.Println("I'm a string:", t)
		default:
			fmt.Println("I'm of a different type:", t)
		}

	}

	whoAmI(42)
	whoAmI("hello")
	whoAmI(3.14)
}
