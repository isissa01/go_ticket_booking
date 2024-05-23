package main

import (
	"fmt"
	"strings"
)

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for their first name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	//ask user for their last name
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	//ask user for their email
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	//validate user input

	//ask user for number of tickets
	fmt.Print("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
