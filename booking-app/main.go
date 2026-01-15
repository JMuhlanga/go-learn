//standard for the main package
package main

import "fmt"

func main(){
	var conferenceName ="Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T ,conferenceName is %T \n",conferenceTickets,remainingTickets,conferenceName)
	// Formatted output - print formatted data - tells GO how to print formatted values, check docs on fmt
	fmt.Printf("Welcome to %v  booking application\n",conferenceName)
	fmt.Printf("We have a total of %v Tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")	
	
	var userName string
	var userTickets int
	// ask user for name
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets \n",userName, userTickets)
}
