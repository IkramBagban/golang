package main

import "fmt"

func main() {
	var users []string = []string{"ikram", "raman"}
	var users2 = make([]string, len(users))

	copy(users2, users) // Copy elements from users to users2 by value
	users2[0] = "ikram2"
	fmt.Print(users)
}
