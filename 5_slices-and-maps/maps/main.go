package main

import "fmt"

func main() {
	var users = make(map[string]string) // Creates a map with string keys and int values
	users["Ikram"] = "Ikram Singh"
	users["raman"] = "Raman Kumar"

	fmt.Print(users)

	fmt.Println("------------")

	m := map[string]string{"name": "bushrah", "surname": "bagban"}

	fmt.Println("m", m)
	delete(m, "surname")
	value, exists := m["surname"]

	if exists {
		fmt.Println("value found", value, exists)
	} else {
		fmt.Println("value not found", value, exists)
	}

}
