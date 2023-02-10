package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	// There are three main types of ways to declare a variable in go

	// The first way
	var firstName string
	firstName = "Hunter"

	// The second way
	lastName := "Pittman"

	// The third way
	var (
		id string
		//totalMoney float32
		time = time.Now()
	)

	fmt.Print("Enter your ID: ")
	fmt.Scanln(&id)

	//fmt.Print

	fmt.Printf("Your ID: %v \n", id)

	fmt.Printf("Your Name: %s \n", firstName+" "+lastName)

	fmt.Printf("The Current Time: %s \n", time)
	goodbye()
}

func goodbye() {
	fmt.Println("Goodbye, World!")
}
