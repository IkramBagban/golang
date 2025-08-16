package main

import "fmt"

func main() {
	var age int = 20
	var name = "Ikram"           // type inferred
	city := "Aurangabad"         // shorthand declaration
	const pi = 3.14              // constant
	isMale := true               // boolean
	const height float32 = 5.12  // float
	var stringNotAssigned string // defaults to ""
	var numberNotAssigned int    // default to 0
	// 	const constNotAssigned int //Constants (const) MUST be assigned a value at declaration time

	fmt.Println(age, name, city, pi, isMale, height, stringNotAssigned, numberNotAssigned)
}
