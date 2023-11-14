package main

import (
	"fmt"
	"strings"
)

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

	for {

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets to buy:")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			bookingDetails = append(bookingDetails, firstName+" "+lastName)
			fmt.Printf("Thankyou %v %v for buying %v tickets. Tickets sent on %v.\n", firstName, lastName, userTickets, email)
	
			firstNames := []string{}
			for _, booking := range bookingDetails {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
	
			fmt.Printf("These are the booking details %v\n", firstNames)
			remainingTickets = remainingTickets - userTickets
	
			if (remainingTickets == 0){
				fmt.Println("Our conference has no tickets left. Come back next year.")
				break
			}
	
			fmt.Printf("Tickets remaining for %v are %v\n", conferenceName, remainingTickets)
		} else {
			fmt.Printf("We only have %v tickets left. %v tickets cannot be booked\n", remainingTickets, userTickets)
		}

	}
}
