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

	conf.GetUserInputs()

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
