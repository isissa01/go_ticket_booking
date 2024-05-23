package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//ask user for their name

	firstName, lastName, email, userTickets := getUserInput()
	//validate user input
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTickets(firstName, lastName, email, userTickets)

		//adds concurrency
		wg.Add(1)

		go sendTicket(firstName, lastName, email, userTickets)

		printFirstNames()

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		printErrorMessages(isValidName, isValidEmail, isValidTicketNumber)
		//validate user input
	}

	wg.Wait()

}
func greetUsers() {
	fmt.Println("Welcome to", conferenceName, "booking system")
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func printErrorMessages(isValidName bool, isValidEmail bool, isValidTicketNumber bool) {
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

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	fmt.Printf("The first names of bookings are: %v\n", firstNames)

}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets -= userTickets
	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)

}
func sendTicket(firstName string, lastName string, email string, userTickets uint) {

	time.Sleep(10 * time.Second)
	fmt.Println("##################")
	fmt.Println("Sending tickets:")
	ticket := fmt.Sprintf("%v tickets for %v %v has been sent to %v\n", userTickets, firstName, lastName, email)

	fmt.Printf("%v", ticket)
	fmt.Println("##################")
	wg.Done()
}
