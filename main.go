package main

import (
	"fmt"

	"mailer/email"
	"mailer/environment"
	"mailer/str"
)

var active bool = true
var lookFor = "atry"

func main() {
	if active {
		config := environment.Get()
		unparsedEmails := email.GetEmails(config.Pop3)
		emails := email.ParseMany(unparsedEmails)
		emailId := str.FindString(emails, lookFor)

		if emailId != -1 {
			response := email.CreateResponse(emails[emailId])

			fmt.Println(response.Body)
			// smtp.SendMail(response)
			active = false
		}

		fmt.Println("Found in email n. ", emailId)

	}
	// check if reply is active
	// if reply is active check new messages
	// if new message contains string we are looking for and sender we are looking for, reply on email, inform that it was replied and set the replier to inactive mode

}
