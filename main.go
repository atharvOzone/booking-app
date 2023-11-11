package main

import "fmt"

func main() {
	var conferenceName = "Atharvs Conference"
	const tickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total tickets: %v Remaining tickets: %v\n", tickets, remainingTickets)
}