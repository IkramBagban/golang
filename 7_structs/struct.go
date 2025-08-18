package main

import (
	"fmt"
	"time"
)

type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

// we can use structs as OOPs in go
type user struct {
	id       int
	name     string
	age      int
	gender   Gender
	createAt time.Time
}

func (u *user) changeName(newName string) {
	u.name = newName
}

func main() {
	user1 := user{
		id:     1,
		name:   "ikram",
		age:    20,
		gender: Male,
	}

	fmt.Println(user1)
	fmt.Printf("%+v\n", user1)

	println("================")
	user1.createAt = time.Now()

	user1.changeName("Urhan")

	fmt.Printf("%+v\n", user1)
}
