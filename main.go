package main
import "fmt"

//Entry point
func main () {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Defining the type
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Printf("Enter your first name:")
	// User input
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address:")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets:")
	fmt.Scan(&userTickets)


	fmt.Printf("Thank you %v %v  for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

		
}
