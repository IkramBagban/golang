package main

import "fmt"

func main() {
	var users []string

	println(users == nil) // true

	var userNames []string = []string{"Ikram", "Bushra"}
	fmt.Println(userNames) // output: [Ikram Bushra]

	println("----- ")
	users2 := make([]string, 3)
	users2[0] = "Ikram"
	users2[1] = "Bushra"
	println("bb ", users2) // output:	 something like this [3/3]0xc000049f00

	println(len(users2))
	println(cap(users2))
	fmt.Print(users2[0] == "")

}
