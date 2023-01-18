package main //Golang is open source, mixture of C and python

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Print("Welcome to ", conferenceName, " booking application \n")
	fmt.Printf("Number of tickets available are %v and the remaining ones are %v \n", conferenceTickets, remainingTickets)

	fmt.Print("*** Get your tickets here to attend the conference *** \n")

	
    
    
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

    fmt.Printf("Thank you %v %v for booking %v tickets. You will get confirmation at %v \n", firstName, lastName, userTicket, email)
	
	remainingTickets = remainingTickets - userTicket
	
    // bookings[0] = firstName + " " + lastName

	//**for appending the array when we dnt knw the size of array** i.e., dynamic list
	bookings = append(bookings, firstName + " " + lastName) 
	
	fmt.Printf("The whole booking array: %v\n", bookings)
    fmt.Printf("The first value booking array: %v \n", bookings[0])
	fmt.Printf("Array type: %T \n", bookings)
	fmt.Printf("Array length: %v \n", len(bookings))
	
	

    

   fmt.Printf("The remaining tickets are %v \n", remainingTickets)

}
