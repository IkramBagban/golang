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

type user struct {
	id       int
	name     string
	age      int
	gender   Gender
	createAt time.Time
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

	fmt.Printf("%+v\n", user1)
}
