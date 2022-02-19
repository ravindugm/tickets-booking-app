package main

import (
	"fmt"
	"tickets-booking-app/conf"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUser(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := conf.GetUserInputs()
		isValidName, isValidEmail, isValidTicketNumber := conf.ValidUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Firstname or Lastname you entered is to short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain '@' sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}

	}

}

func greetUser(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Println("--------------------------------------------------")
	fmt.Println("|      Welcome to RMDev Tickets Booking App      |")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("**************************************************")
	fmt.Println("**************************************************")
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}
