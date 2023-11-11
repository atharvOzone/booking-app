package main

import "fmt"

func main() {
	conferenceName := "Atharvs Conference"
	const tickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total tickets: %v Remaining tickets: %v\n", tickets, remainingTickets)

	var userName string
	var userTickets uint

	userName="Golang"
	userTickets=50

	fmt.Printf("User %v booked %v tickes", userName, userTickets)
}