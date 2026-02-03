//standard for the main package
package main

import (
	"fmt"
	"strings"
)

func main(){
	var conferenceName string ="Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings = []string{}
	
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T ,conferenceName is %T \n",conferenceTickets,remainingTickets,conferenceName)
	// Formatted output - print formatted data - tells GO how to print formatted values, check docs on fmt
	fmt.Printf("Welcome to %v  booking application\n",conferenceName)
	fmt.Printf("We have a total of %v Tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")	
	
	for remainingTickets > 0 && len(bookings) < conferenceTickets {		
		var firstName string
		var lastName string
		var email string
		var userTickets int
		var remainingTickets int
		
		// ask user for name
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)
		
		var isValidName bool = len(firstName) >=2 && len(lastName) >=2
		var isValidEmail bool = strings.Contains(email,"@")
		var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
		
			if userTickets < remainingTickets {
				remainingTickets = remainingTickets - int(userTickets)
	
				bookings = append(bookings, firstName +" "+lastName)
				
				fmt.Printf("The Whole slice is %v\n", bookings)
				fmt.Printf("The first element is %v\n", bookings[0])
				fmt.Printf("The slice type is %T\n", bookings)
				fmt.Printf("The slice length is %v\n", len(bookings))
				
				fmt.Printf("User %v %v booked %v tickets you will get an email at %v\n",firstName,lastName, userTickets,email)
				fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)
				
				firstNames := []string{}
				for _, booking := range bookings {
					var names = strings.Fields(booking)
					var firstName = names[0]
					firstNames = append(firstNames, firstName)
				}
				fmt.Printf("These first names of bookings are: %v\n", firstNames)
				
				var noTicketsRemaining bool = remainingTickets == 0
				if noTicketsRemaining {
					fmt.Println("Our conference is booked out. Come back next year.")
					break
				}
			}else if userTickets == remainingTickets {
				//Do something else
				fmt.Println("You have booked all the tickets")
				break
			}else{
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			}		
		}else{
			if !isValidName {
				fmt.Println("Invalid name, please try again, first name or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Invalid email, please try again, Email address you entered is invalid as it does not contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticket number, please try again, number of tickets entered is invalid")
			}			
		}	
}
