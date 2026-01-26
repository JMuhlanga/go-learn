//standard for the main package
package main

import "fmt"

func main(){
	var conferenceName string ="Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T ,conferenceName is %T \n",conferenceTickets,remainingTickets,conferenceName)
	// Formatted output - print formatted data - tells GO how to print formatted values, check docs on fmt
	fmt.Printf("Welcome to %v  booking application\n",conferenceName)
	fmt.Printf("We have a total of %v Tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")	
	
	var bookings = []string{}
	
	var firstName string
	var lastName string
	var email string
	var userTickets int
		
	// ask user for name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	
	remainingTickets = remainingTickets - uint(userTickets)

	bookings = append(bookings, firstName +" "+lastName)
	
	fmt.Printf("The Whole slice is %v\n", bookings)
	fmt.Printf("The first element is %v\n", bookings[0])
	fmt.Printf("The slice type is %T\n", bookings)
	fmt.Printf("The slice length is %v\n", len(bookings))
	
	fmt.Printf("User %v %v booked %v tickets you will get an email at %v\n",firstName,lastName, userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)
}
