package main

import (
	"fmt"
	"sync"
	"time"
)

//fmt or format is a built in go package from where print comes

const conferenceTicket int = 50

// var conferenceName = "Go conference"
var conferenceName = "Go conference"
var remainingTickets uint = 50

// var bookings [50]string // array
// dynamic array -> slice
// var bookings []string
// var bookings = make([]map[string]string, 0) // slice of maps
var bookings = make([]userData, 0)

// struct is a lightweight class(predefined structure) without having inheritence
// custom-type
type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// "Waitgroup1" waits for launched goroutine to finish
// sync provides basic synchronization funcitonality
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTicekts := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicekts)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTickets(firstName, lastName, email, userTicekts)

		//Add-> set the number of goroutines to wait for(increase the counter by providing the number)
		wg.Add(1)

		// "go..." starts a new goroutine(a lightweight thread managed by the go runtime)
		// this keword abstract away the whole thread creation and ensure concurrency
		go sendTicket(userTicekts, firstName, lastName, email)

		// fmt.Printf("The first names of bookings are: %v\n", getFirstNames(bookings))
		firstNames := getFirstNames()

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year!!")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}

	//Wait func waits for all the threads added before are done or
	// blocks untill the WaitGroupt counter is 0
	wg.Wait()
}

func greetUsers() {
	//Printf is called print format funciton allows to format the variables
	fmt.Printf("Wellcome to %v booking application\n", conferenceName) //%v is a placeholder
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	// blank indentifier--> to igonre a variable that we don't want to use
	// underscore(_) used to tell go index will not be used in the program
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicekts uint

	fmt.Printf("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address:")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets:")
	fmt.Scan(&userTicekts)
	return firstName, lastName, email, userTicekts
}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	// calculaiton must be performed between same types
	remainingTickets = remainingTickets - userTickets
	// // bookings[0] = firstName + " " + lastName
	// bookings = append(bookings, firstName+" "+lastName)

	//creating a empty map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	// bookings = append(bookings, userData)

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickts. You will get an confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

// generating a ticket after booking and send it via email
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// sleep function stops or blocks the current thread or goroutine execution for 10s
	time.Sleep(10 * time.Second)
	// Sprint fun return a string rather printing
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket :\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")

	// Decrements the WaitGroup counter by 1
	// this is called by the goroutine to indicate that it's finished
	// it removes the thread that was added beforef
	wg.Done()
}

// to run the all files within the application -> go run .
// to quit the application from CLI cntrl + c
// use go run .(folder location from which all files will be excuted)
// when there are multiple files to run within the same package
