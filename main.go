package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	var bookings []string

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	//ask user for their name
	for {
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

		//validate user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets -= userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)

			printFirstNames(bookings)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			validateUserData(isValidName, isValidEmail, isValidTicketNumber)
			//validate user input
		}
	}

}
func greetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Println("Welcome to", conferenceName, "booking system")
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func validateUserData(isValidName bool, isValidEmail bool, isValidTicketNumber bool) {
	if !isValidName {
		fmt.Println("Name must be at least 2 characters")

	}
	if !isValidEmail {
		fmt.Println("Email address is invalid")

	}
	if !isValidTicketNumber {
		fmt.Println("Number of tickets is invalid")
	}
}

func printFirstNames(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, strings.Split(booking, " ")[0])
	}

	fmt.Printf("The first names of bookings are: %v\n", firstNames)

}
