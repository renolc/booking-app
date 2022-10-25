package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceName = "Go Conference"
const conferenceTickets = 50

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

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets)

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

		userData := bookTickets(firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendTicket(userData)

		printFirstNames()
	}

	fmt.Println("Our conference is booked out. Come back next year.")
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	fmt.Printf("These first names of our bookings: %v\n", firstNames)
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName, lastName, email string, userTickets uint) UserData {
	remainingTickets -= userTickets

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Current bookings: %v\n", bookings)

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %s\n", remainingTickets, conferenceName)
	return userData
}

func sendTicket(userData UserData) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%d tickets for %s %s", userData.numberOfTickets, userData.firstName, userData.lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %s\nto email %s\n", ticket, userData.email)
	fmt.Println("######################")
	wg.Done()
}
