package main

import "fmt"

func main() {
	// pass by reference
	var users = make(map[string]string) // Creates a map with string keys and string values
	users["Ikram"] = "Ikram bagban"
	users["raman"] = "Raman bagban"
	users2 := users
	users2["Ikram"] = "Ikram bagban 2"
	fmt.Print(users)

}
