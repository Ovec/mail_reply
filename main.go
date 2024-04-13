package main

import (
	"mailer/environment"
	"mailer/pop"
)

var active bool = true

func main() {
	if active {
		config := environment.Get()
		pop.GetEmails(config.Pop3)
	}
	// check if reply is active
	// if reply is active check new messages
	// if new message contains string we are looking for and sender we are looking for, reply on email, inform that it was replied and set the replier to inactive mode

}
