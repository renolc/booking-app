package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets

		if !isValidName || !isValidEmail || !isValidTicketCount {
			if !isValidName {
				fmt.Println("First or last name entered are too short")
			}
			if !isValidEmail {
				fmt.Println("Email must contain @ symbol")
			}
			if !isValidTicketCount {
				fmt.Println("Number of tickets is invalid")
			}
			continue
		}

		remainingTickets -= userTickets
		bookings = append(bookings, fmt.Sprintf("%s %s", firstName, lastName))

		fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
		fmt.Printf("%d tickets remaining for %s\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			firstNames = append(firstNames, strings.Fields(booking)[0])
		}

		fmt.Printf("These first names of our bookings: %v\n", firstNames)
	}

	fmt.Println("Our conference is booked out. Come back next year.")
}
