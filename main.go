package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Println("Welcome to", conferenceName, "booking system")
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	var bookings []string

	//ask user for their name
	for {

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
		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			return
		}

		remainingTickets -= userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			firstNames = append(firstNames, strings.Split(booking, " ")[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}

	}

	//send ticket to user email

	//call function to print out summary

	//call function to send ticket to email

}
