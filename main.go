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
	var userName string
	var userTickets int

	// User input
	fmt.Scan(userName)

	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

		
}
