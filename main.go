package main

import (
	"fmt"
	"os"
)

func main() {

	var eventName = "CISO Super Cyber Summit"
	const totalTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to the %v! \n", eventName)
	fmt.Printf("There are %v tickets remainning at %s \n", remainingTickets, eventName)

	// & creates a pointer (sometimes called a borrow) and * gets the value of a pointer (also called a dereference)

	randomValue1 := "bruh"
	var memoryAddressOfRandomValue1 *string
	memoryAddressOfRandomValue1 = &randomValue1

	fmt.Printf("The memory address of randomValue1 is %v \n", memoryAddressOfRandomValue1)

	infoGathering(totalTickets, remainingTickets)

}

func infoGathering(totalTickets int, remainingTickets int) {
	var (
		id          string
		firstName   string
		lastName    string
		email       string
		userTickets int
	)

	fmt.Print("Enter your ID: ")
	fmt.Scan(&id)

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)

	if userTickets > remainingTickets {
		fmt.Println("I am sorry there are not enough tickets left to complete your transaction. Please try again later. Ending program!")
		os.Exit(0)
	} else if userTickets == remainingTickets {
		fmt.Println("Congratualtions! You have purchased the last ticket for this event! Ending program!")
		os.Exit(0)
	} else {
		remainingTickets -= userTickets
		fmt.Printf("There are %v tickets left for this event. \n", remainingTickets)
	}

	fmt.Println("###########Transaction Information###########")
	fmt.Printf("Your ID: %v \n", id)
	fmt.Printf("Your Name: %s \n", firstName+" "+lastName)
	fmt.Printf("Your email: %s \n", email)
	fmt.Printf("Number of tickets purchased: %v \n", userTickets)
}
