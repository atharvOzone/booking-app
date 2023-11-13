package main

import "fmt"

func main() {
	conferenceName := "Atharvs Conference"
	const tickets int = 50
	var remainingTickets uint = 50
	var bookingDetails []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total tickets: %v Remaining tickets: %v\n", tickets, remainingTickets)

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

	bookingDetails = append(bookingDetails, firstName + " " + lastName)
	fmt.Printf("Thankyou %v %v for buying %v tickets. Tickets sent on %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("These are the booking details %v\n", bookingDetails)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Tickets remaining for %v are %v\n", conferenceName, remainingTickets)
}