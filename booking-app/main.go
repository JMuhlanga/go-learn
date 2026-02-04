package main

import (
	"fmt"
	"booking-app/helper"
	//"strconv"
	// "strings"
	"time"
	"sync"
)

var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var remainingTickets int = 50
var bookings = make([]UserData,0) 

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets int
	isOptedInForNewsletter bool
}
var wg = sync.WaitGroup{}

func main() {
	// Greet Users
	greetUsers()

	//for remainingTickets > 0 && len(bookings) < conferenceTickets {
		//Get user input
		firstName,lastName,email,userTickets := getUserInput()
		// Validation
		isValidName,isValidEmail,isValidTicketNumber := helper.ValidateUserInput(firstName,lastName,email,userTickets,remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Book Ticket
			bookTicket(userTickets, firstName, lastName, email)
			// Send Ticket
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			
			// Print first names
			var firstNames = printFirstNames()
			fmt.Printf("The first names of bookings are: %v\n",firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
			//	break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name: first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Invalid email: email must contain @")
			}
			if !isValidTicketNumber {
				fmt.Printf(
					"Invalid ticket number: we only have %v tickets remaining\n",
					remainingTickets,
				)
			}
		}
		//}
	
	// city := "London"
	
	// switch city{
	// 	case "New York":
	// 		fmt.Println("You are booking tickets for New York")
	// 	case "Singapore","Hong Kong":
	// 		fmt.Println("You are booking tickets for Singapore")
	// 	case "Berlin","Lodon":
	// 		fmt.Println("You are booking tickets for Berlin")
	// 	case "Mexico City":
	// 		fmt.Println("You are booking tickets for Mexico City")
	// 	default:
	// 		fmt.Println("No Valid City Selected")
	// }
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {		
		//firstNames = append(firstNames, booking["firstName"] )
		firstNames = append(firstNames, booking.firstName )
	}
	return firstNames
}

func getUserInput()(string,string,string,int){
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// Ask user for details
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	
	return firstName,lastName,email,userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	//create Map for user 
	//var userData = make(map[string]string)
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"]=strconv.FormatInt(int64(userTickets),10)
	
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n",bookings)

	fmt.Printf("The whole slice is %v\n", bookings)
	fmt.Printf("The first element is %v\n", bookings[0])
	fmt.Printf("The slice type is %T\n", bookings)
	fmt.Printf("The slice length is %v\n", len(bookings))

	fmt.Printf(
		"User %v %v booked %v tickets. You will get an email at %v\n",
		firstName, lastName, userTickets, email,
	)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName string, lastName string,email string){
	//Go Routine implementation
	time.Sleep(10 * time.Second) // stops execution of thread for 10 seconds
	
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("##########\n")
	fmt.Printf("Sending ticket:\n %v to email address %v\n",ticket,email)
	fmt.Printf("##########\n")
	
	wg.Done()
}