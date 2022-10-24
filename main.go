package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
}
