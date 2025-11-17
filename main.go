package main

import (
	"fmt"
	"strings"
)

func main() {

	bookings := []string{} // slice, dynamic in size
	const totalTickets uint = 50
	var firstName, lastName, email string
	var tickets uint
	var remaining uint = totalTickets
	for {

		fmt.Println("--------------- Welcome to Air Tarsai ---------------")
		fmt.Println("Plz Locked in your Tickets Because You are Gonna Achieve the Heaven With us!")

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets you want: ")
		fmt.Scan(&tickets)

		if tickets > remaining {
			fmt.Printf("Sorry, only %v tickets remaining.\n", remaining)
			return
		}

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := tickets > 0 && tickets <= remaining

		if isValidEmail && isValidName && isValidTicketNumber {
			remaining -= tickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\nBooking confirmed for %v %v!\n", firstName, lastName)
			fmt.Printf("You booked %v tickets. %v remaining.\n", tickets, remaining)
			firstnames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)

				firstnames = append(firstnames, names[0])
			}
			fmt.Printf("Current bookings: %v\n", firstnames)

			if remaining == 0 {
				fmt.Println("Sorry! But We Have No Tickets For You , Try Again Next Time!!")
			}
		} else {
			if !isValidName {
				fmt.Println("Your name is not a valid, Your name must be contain more than or equal to 2 words!")
			} else if !isValidEmail {
				fmt.Println("Your email is not a valid!")
			} else {
				fmt.Println("Ticket Numbers Are Not valid")
			}
		}

	}
}
