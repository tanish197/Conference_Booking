package main //Golang is open source, mixture of C and python

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Print("Welcome to ", conferenceName, " booking application \n")
	fmt.Printf("Number of tickets available are %v and the remaining ones are %v \n", conferenceTickets, remainingTickets)

	fmt.Print("*** Get your tickets here to attend the conference *** \n")

	for {
		var bookings = []string{}
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		println("Enter your first name")
		fmt.Scan(&firstName) // & it is used for pointer in Golang

		println("Enter your last name")
		fmt.Scan(&lastName)

		println("Enter your email")
		fmt.Scan(&email)

		println("Enter number of tickets you want to book")
		fmt.Scan(&userTicket)

		isValidName := len(firstName) >= 2 &&  len(lastName) >= 2
        isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTicket >0 && userTicket <= remainingTickets
		
		if isValidName && isValidEmail && isValidTicketNumber{
        if userTicket > remainingTickets{
			fmt.Printf("We only have %v tickets available, so you can't book %v tickets\n ", remainingTickets, userTicket)
			break // loop will break here
			// continue  **This will not break and jump to next iteration in case user wants to correct number of tickets** 
		}
		remainingTickets = remainingTickets - userTicket

		// bookings[0] = firstName + " " + lastName

		//**for appending the array when we dnt knw the size of array** i.e., dynamic list
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("Thank you %v %v for booking %v tickets. You will get confirmation at %v \n", firstName, lastName, userTicket, email)
		fmt.Printf("%v tickets remaining for GO conference\n", remainingTickets)

		// fmt.Printf("The whole booking array: %v\n", bookings)
		// fmt.Printf("The first value booking array: %v \n", bookings[0])
		// fmt.Printf("Array type: %T \n", bookings)
		// fmt.Printf("Array length: %v \n", len(bookings))
		// fmt.Printf("Slice type: %T \n", bookings)
		// fmt.Printf("Slice length: %v \n", len(bookings))

		// fmt.Printf("These all are booking: %v\n", bookings)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our Conference is booked out. Come back next year")
			break

		}
	} else {
		fmt.Println("Your input data is invalid, try again")
	}
} 

}
