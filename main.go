package main

import (
	"booking-app/helpers"
	"fmt"
)

var conferenceName = "Atharvs Conference"

const tickets int = 50

var remainingTickets uint = 50
var bookingDetails = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTickets := helpers.ValidateInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTickets {
			bookTickets(firstName, lastName, userTickets, email)
			firstNames := printFirstNames()
			fmt.Printf("These are the booking details %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference has no tickets left. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name and last name are too short!!")
			}
			if !isValidEmail {
				fmt.Println("Invalid Email")
			}
			if !isValidTickets {
				fmt.Println("Invalid number of tickets, kindly enter >0 or <=50")
			}
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total tickets: %v Remaining tickets: %v\n", tickets, remainingTickets)
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookingDetails {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets to buy:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint, email string) {
	//create a user map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookingDetails = append(bookingDetails, userData)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thankyou %v %v for buying %v tickets. Tickets sent on %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("Tickets remaining for %v are %v\n", conferenceName, remainingTickets)
}
