package main

import "fmt"

func main() {
	var conferenceName = "Atharvs Conference"
	const tickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("Total tickets:", tickets, "Remaining tickets:", remainingTickets)
}