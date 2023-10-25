package main

import "fmt"
import "strings"
// a ticket booking application

func main(){
   var name = "james-conference"
   var conferenceTicket = 10
   var remainingTicket uint = 10

   for {
   fmt.Println("hello world")
   fmt.Println("Welcome to", name, "We are excited to host the life-changing conference")
   fmt.Println("Please, kindly book a ticket to attend the conference. We have", conferenceTicket, "tickets available for the conference")

   // let us collect user data as they book the tickeks
   // this will actually put the input on one line like in pyhton

   var booker_name string
   var booker_email string
   var booker_no_of_tickets uint 
   var isValidbooker_name = len(booker_name) >= 2
   var isValidbooker_email = strings.Contains(booker_email, "@")
   var isvalidbooker_no_of_tickets = booker_no_of_tickets > 0

   fmt.Print("Please enter your name: ")
   fmt.Scan(&booker_name)
   fmt.Print("Enter your email address: ")
   fmt.Scan(&booker_email)
   fmt.Print("How many tickets do you wish to book?: ")
   fmt.Scan(&booker_no_of_tickets)

   remainingTicket = remainingTicket - booker_no_of_tickets

   if remainingTicket == 0{
      // end program
      fmt.Println("Sorry, all tickets have been sold out, please come back next year")
      break
   } else {
      if booker_no_of_tickets > remainingTicket{
         fmt.Println("Sorry, we only have", conferenceTicket, "so you can't book", booker_no_of_tickets, "tickets")
         continue
      } else {
         if !isValidbooker_name{
            fmt.Println("The value for name must be more than 2 characters")
            continue
         }
         if !isValidbooker_email{
            fmt.Println("The value for email must have '@' synbol")
            continue
         }
         if !isvalidbooker_no_of_tickets{
            fmt.Println("You can't book 0 number of tickets")
            continue
         }
         fmt.Println("Congratulations", booker_name, "you have successfully booked", booker_no_of_tickets, "tickets")
         fmt.Println("You will be sent a confirmation mail at", booker_email,)
      }
   }

   }
   
}