package email

import (
	"fmt"
	"io"
	"mime/quotedprintable"

	"github.com/emersion/go-message"
)

type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

func ParseMany(messages []*message.Entity) []*Email {
	var emails []*Email

	fmt.Println("Printing from parse many")
	fmt.Println(len(messages))

	for _, m := range messages {
		message, _ := io.ReadAll(quotedprintable.NewReader(m.Body))

		emails = append(emails, &Email{
			From:    m.Header.Get("from"),
			To:      m.Header.Get("to"),
			Subject: m.Header.Get("subject"),
			Body:    string(message),
		})
	}

	return emails
}

func CreateResponse(email *Email) *Email {
	response := &Email{
		From:    email.To,
		To:      email.From,
		Subject: email.Subject,
		Body:    "Jedu.\n\n -----Original Message-----\n" + email.Body + "\n\n",
	}

	return response
}
