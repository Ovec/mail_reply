package email

import (
	"fmt"
	"io"
	"mailer/environment"
	"mime/quotedprintable"

	"github.com/emersion/go-message"
	"github.com/google/uuid"
)

type Email struct {
	Headers message.Header
	// From    string
	// To      string
	// Subject string
	Body string
}

func ParseMany(messages []*message.Entity) []*Email {
	var emails []*Email

	for _, m := range messages {
		message, _ := io.ReadAll(quotedprintable.NewReader(m.Body))

		emails = append(emails, &Email{
			Headers: m.Header,
			// From:    m.Header.Get("from"),
			// To:      m.Header.Get("to"),
			// Subject: m.Header.Get("subject"),
			Body: string(message),
		})
	}

	return emails
}

func CreateResponse(email *message.Entity, config environment.Config) *message.Entity {
	printMail(email)
	fmt.Println()

	// create new headers
	headers := message.Header{}

	messageId := "<" + uuid.New().String() + "@" + config.Smtp.Host + ">"

	headers.Set("from", "<"+config.Email.Sender+">")
	// email.Headers.Set("To", "mirekovec@gmail.com")
	headers.Set("To", email.Headers.Get("from"))
	headers.Set("Subject", "Re: "+email.Headers.Get("subject"))
	headers.Set("References", email.Headers.Get("Message-Id")+" "+email.Headers.Get("References"))
	headers.Set("In-Reply-To", email.Headers.Get("Message-Id"))
	headers.Set("Message-Id", messageId)
	headers.Set("Content-Type", "text/plain; charset=UTF-8")
	headers.Set("Content-Transfer-Encoding", "quoted-printable")

	response := &Email{
		Headers: headers,
		Body:    string(email.Body),
	}

	printMail(response)

	// response := &Email{
	// 	From:    email.To,
	// 	To:      email.From,
	// 	Subject: "Re:" + email.Subject,
	// 	Body:    "Jedu.\n\n -----Original Message-----\n" + email.Body + "\n\n",
	// }

	return response
}

// Function to print an Entity's details
func printEntityDetails(entity *message.Entity) {
	// Print the media type and media parameters
	fmt.Printf("Media Type: %s\n", entity.mediaType)

	fmt.Println("Media Parameters:")
	for key, value := range entity.mediaParams {
		fmt.Printf("  %s: %s\n", key, value)
	}

	// Print the header
	fmt.Println("Header:")
	fmt.Print(entity.Header.String()) // Custom String method to output the header

	// Print the body (be careful to avoid side effects with io.Reader)
	fmt.Println("Body:")
	bodyContent, err := io.ReadAll(entity.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	} else {
		fmt.Println(string(bodyContent))
	}
}
