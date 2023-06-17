package main

import (
	"fmt"
	"strings"
	//"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
)

func main() {
	//a := app.New()
	//w := a.NewWindow("Hello World")
	//w.Resize(fyne.NewSize(300, 300))
	//w.ShowAndRun()
	//w.SetContent(

	var conferencename = "go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("welcome to our %v booking app.\n", conferencename)
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")

	for {

		var firstname string
		var lastname string
		var email string
		var userTickets uint

		isValidName := len(firstname) >= 2 && len(lastname) >= 2
		isvalidemail := strings.Contains(email, "@")
		isvalidticketnumber := userTickets > 0 && userTickets <= remainingTickets

		// ask user for their name
		fmt.Println("please enter your first name:")
		fmt.Scanln(&firstname)
		if !isValidName {
			fmt.Println("invalid first name or last name")
		}

		// ask user for their last name
		fmt.Println("please enter your last name:")
		fmt.Scanln(&lastname)

		// ask user for their email
		fmt.Println("please enter your email address:")
		fmt.Scanln(&email)

		// ask user for their tickets
		fmt.Println("please enter number of tickets:")
		fmt.Scanln(&userTickets)

		//if isValidName && isvalidemail && isvalidticketnumber {

		if userTickets < 1 {
			fmt.Println("you must enter at least one ticket")
		}
		if userTickets > conferenceTickets {
			fmt.Println("you can't have more tickets than the available tickets")
			continue
		}
		if userTickets > remainingTickets {
			fmt.Println("you can't have more tickets than the remaining tickets")
			continue
		}
		if firstname == "" || lastname == "" || email == "" {
			fmt.Println("you must enter at least one name, lastname and email")
			continue
		}
		//isValidName := len(firstname) >= 2 && len(lastname) >= 2
		//isvalidemail := strings.Contains(email, "@")
		//isvalidticketnumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isvalidemail && isvalidticketnumber {
			fmt.Printf("your name is %v %v \n", firstname, lastname)
			fmt.Printf("your tickets are %v\n", userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstname+" "+lastname+" "+email)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstname, lastname, userTickets, email)
		fmt.Printf("%v tickets are available for booking %v.\n", remainingTickets, conferencename)

		firstnames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstnames = append(firstnames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstnames)

		if remainingTickets == 0 {
			//end of app
			fmt.Println("Our conference is booked out, come back next year, Thanks for your booking!")
			break

		}
		var answer string
		fmt.Println("Do you want to continue? (y/n)")

		if strings.ToLower(answer) == "y" {
			continue
		}
		if strings.ToLower(answer) == "n" {
			break
		}
		fmt.Println("please enter y or n")
		fmt.Scanln(&answer)
		if strings.ToLower(answer) == "y" {
			continue
		}
		if strings.ToLower(answer) == "n" {
			break
		}

		//}

	}
}
